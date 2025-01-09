package mongo

import (
	"cre/styles"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

// CreateCredentials generates MongoDB credentials and saves them to files.
// It takes a pointer to a MongoCredentials struct and an output directory path.
// It returns an error if the output directory cannot be created or if the password
// generation or file writing fails.
func CreateCredentials(c *MongoCredentials, output string) error {

	// check if exsist the output folder. If not, create it
	if _, err := os.Stat(output); os.IsNotExist(err) {
		log.Info("Creating output directory")
		if err := os.Mkdir(output, 0755); err != nil {
			return err
		}
	}

	log.Info("Generating credentials...")

	// se è un utente ROOT, assegna il nome root all'username e il database admin
	if c.Root {
		log.Warn("User is a ROOT user")
		c.setRoot()
	}

	// Convert all string fields to lowercase
	c.toLower()

	// creating the mongoUserPasswordFile as a pattern MONGO_<USERNAME>_PW.
	mongoPasswordFilename := fmt.Sprintf("%s/MONGO_%s_PW", output, strings.ToUpper(c.Username))

	// Check if the file already exists
	if _, err := os.Stat(mongoPasswordFilename); err == nil {
		log.Warn("File already exists\n")

		// Prompt the user to confirm overwriting the file
		if overwritePW := overwriteConfim(); overwritePW {
			passwordGeneratorLogger(c, mongoPasswordFilename)
		} else {
			log.Warn("Skip password generation. Keeping the existing credentials")
			if err := c.readMongoPasswordFromFile(mongoPasswordFilename); err != nil {
				return err
			}
		}
	} else {
		err := passwordGeneratorLogger(c, mongoPasswordFilename)
		if err != nil {
			return err
		}
	}

	if r, _ := isReplicaSet(c.Host); r {
		log.Info("Replica Set detected")
	} else {
		log.Info("Single host detected")
	}

	log.Info("Compiling connection string...")
	c.compileConnectionString()

	// Write a resume of the credentials in a yaml file called "mongo_<username>_credentials.yaml
	log.Info("Writing credential file...")
	credentialFilename := fmt.Sprintf("%s/mongo_%s_credentials.yaml", output, strings.ToLower(c.Username))

	if yamlBytes, err := c.saveCredentialsAsYaml(credentialFilename); err != nil {
		return err
	} else {
		log.Info(fmt.Sprintf("Credentials written %d bytes to %s", yamlBytes, credentialFilename))
	}
	return nil
}

// log the password generation process while writing the password to a file
func passwordGeneratorLogger(c *MongoCredentials, mongoPasswordFilename string) error {
	log.Info("Generating password...")
	if pwdBytes, err := c.createPasswordFile(mongoPasswordFilename); err != nil {
		return err
	} else {
		log.Info(fmt.Sprintf("Wrote %d bytes to file %v", pwdBytes, mongoPasswordFilename))
		return nil
	}
}


func CredentialsResume(c *MongoCredentials) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Root: %t\n", (c.Root))
	fmt.Fprintf(&sb, "Username: %s\n", styles.PaintRed(c.Username))
	fmt.Fprintf(&sb, "Database: %s\n", styles.PaintCyan(c.Database))
	fmt.Fprintf(&sb, "Host: %s\n", styles.PaintBlue(c.Host))
	fmt.Fprintf(&sb, "Replica Set: %v\n", styles.PaintYellow(c.ReplicaSet))
	fmt.Fprintf(&sb, "Password: %s\n", styles.PaintGray(c.Password))
	fmt.Fprintf(&sb, "Authentication DB: %s\n", styles.PaintOrange(c.AuthenticationDB))
	return sb.String()
}