package mongo

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"crypto/rand"

	yaml "gopkg.in/yaml.v3"
)

type MongoCredentials struct {
	Username         string
	Host             string
	ReplicaSet       string
	Database         string
	Password         string
	Root             bool
	ConnectionString string
	AuthenticationDB string
}

// compileConnectionString constructs the MongoDB connection string based on the
// MongoCredentials fields and assigns it to the ConnectionString field. It includes
// the username, password, host, database, and authentication database. If the host
// is part of a replica set, it appends the replica set information to the connection
// string. The method returns the updated MongoCredentials instance.
func (c *MongoCredentials) compileConnectionString() *MongoCredentials {
	var sb strings.Builder
	fmt.Fprintf(&sb, "mongodb://%s:%s@%s/%s?authSource=%s",
		c.Username,
		c.Password,
		c.Host,
		c.Database,
		c.AuthenticationDB,
	)
	if r, _ := isReplicaSet(c.Host); r {
		fmt.Fprintf(&sb, "&replicaSet=%s", c.ReplicaSet)
	}
	c.ConnectionString = sb.String()
	return c
}

// setRoot sets the MongoCredentials to use root credentials.
// It sets the Root field to true, the Username to "root", and the Database to "admin".
// Returns the updated MongoCredentials instance.
func (c *MongoCredentials) setRoot() *MongoCredentials {
	c.Root = true
	c.Username = "root"
	c.Database = "admin"
	return c
}

// toLower converts all string fields of the MongoCredentials struct to lowercase.
// This includes Username, Database, Host, ReplicaSet, and AuthenticationDB fields.
// It returns a pointer to the modified MongoCredentials struct.
func (c *MongoCredentials) toLower() *MongoCredentials {
	c.Username = strings.ToLower(c.Username)
	c.Database = strings.ToLower(c.Database)
	c.Host = strings.ToLower(c.Host)
	c.ReplicaSet = strings.ToLower(c.ReplicaSet)
	c.AuthenticationDB = strings.ToLower(c.AuthenticationDB)
	return c
}

// createPasswordFile generates a random password and writes it to a file named
// after the MongoDB username. The password is saved to the Password field of the
// MongoCredentials struct. It returns the number of bytes written to the file and
// an error if the password generation or file writing fails.
func (c *MongoCredentials) createPasswordFile(mongoPasswordFilename string) (int, error) {
	password, err := generateURLSafePassword(64)
	if err != nil {
		e := errors.New("dag-error generating password")
		errors.Join(e, err)
		return 0, e
	} else {
		c.Password = password
	}
	// Write the credentials to a file
	file, err := os.Create(mongoPasswordFilename)
	if err != nil {
		e := errors.New("dag-error creating password file")
		errors.Join(e, err)
		return 0, e
	}
	pwdBytes, err := file.WriteString(c.Password)
	if err != nil {
		e := errors.New("dag-error writing password file")
		errors.Join(e, err)
		return 0, e
	}
	defer file.Close()
	return pwdBytes, nil
}

// readMongoPasswordFromFile reads the password from a file and saves it to the
// Password field of the MongoCredentials struct. It returns an error if the file
// cannot be opened or read.
func (c *MongoCredentials) readMongoPasswordFromFile(mongoPasswordFilename string) error {
	// read the password from the file and save it to the credential struct
	file, err := os.Open(mongoPasswordFilename)
	if err != nil {
		e := errors.New("dag-error opening password file")
		errors.Join(e, err)
		return e
	}
	defer file.Close()

	// read the password string from the file
	content, err := io.ReadAll(file)
	if err != nil {
		e := errors.New("dag-error reading password file")
		errors.Join(e, err)
		return e
	}
	c.Password = string(content)
	return nil
}

// saveCredentialsAsYaml saves the MongoCredentials struct to a YAML file with the
// provided filename. It returns the number of bytes written to the file and an
// error if the file cannot be created or written to.
func (c *MongoCredentials) saveCredentialsAsYaml(credentialFilename string) (int, error) {
	credFile, err := os.Create(credentialFilename)
	if err != nil {
		e := errors.New("dag-error creating credential file")
		errors.Join(e, err)
		return 0, e
	}

	// Transform the MongoCredentials struct into a YAML file
	yamlData, err := yaml.Marshal(c)
	if err != nil {
		e := errors.New("dag-error marshaling credentials to YAML")
		errors.Join(e, err)
		return 0, e
	}

	// Write the YAML data to the file
	b, err := credFile.Write(yamlData)
	if err != nil {
		e := errors.New("dag-error writing YAML file")
		errors.Join(e, err)
		return 0, e
	}

	return b, nil
}

////////////////////////////////////////////////////////////////////////////////
// Helper functions
////////////////////////////////////////////////////////////////////////////////

// validateHostname checks if the provided hostname matches the expected pattern.
// The expected pattern is a string containing a hostname followed
// by a colon and a port number (e.g., "example.com:1234").
//
// It returns a boolean indicating whether the hostname is valid and an error
// if the regex compilation fails.
//
// Parameters:
//   - hostname: The hostname string to validate.
//
// Returns:
//   - bool: True if the hostname is valid, false otherwise.
//   - error: An error if the regex compilation fails, nil otherwise.
func validateHostname(hostname string) (bool, error) {
	pattern, err := regexp.Compile(`[a-z0-9\.\-]+:\d+`)
	matches := pattern.FindAllString(hostname, -1)
	return err != nil || matches == nil, err
}

// isReplicaSet checks if the given hostname belongs to a MongoDB replica set.
// A hostname is considered part of a replica set if it contains more than one
// host:port pair.
//
// Parameters:
//   - hostname: A string representing the MongoDB hostname.
//
// Returns:
//   - bool: True if the hostname is part of a replica set, false otherwise.
//   - error: An error if the regular expression compilation fails.
func isReplicaSet(hostname string) (bool, error) {
	pattern, err := regexp.Compile(`[a-z0-9\.\-]+:\d+`)
	matches := pattern.FindAllString(hostname, -1)
	return len(matches) > 1, err
}

func generateURLSafePassword(length int) (string, error) {
	// Calculate the number of random bytes needed to generate a base64 string of the desired length
	// Since base64 encodes 3 bytes into 4 characters, we need length * 3/4 bytes
	byteLength := (length * 3) / 4

	// Generate random bytes
	randBytes := make([]byte, byteLength)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}

	// Encode the bytes to a URL-safe base64 string
	password := base64.URLEncoding.EncodeToString(randBytes)

	// Truncate the string to the desired length (base64 encoding may produce a slightly longer string)
	if len(password) > length {
		password = password[:length]
	}

	return password, nil
}
