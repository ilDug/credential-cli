package mongo

import (
	"fmt"
	"os"
	"strings"

	"cre/styles"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// Run the program and the form
func MongoRun(outDir ...string) {
	var output string
	if len(outDir) == 0 || outDir[0] == "" {
		output = "./secret"
	} else {
		output = outDir[0]
	}

	var credentials = MongoCredentials{}
	credentials.AuthenticationDB = "admin"
	credentials.ReplicaSet = ""

	form := formInit(&credentials)

	// Run the form.
	err := form.Run()

	// Handle form errors.
	if err != nil {
		// If the user aborted the form, exit gracefully.
		if err == huh.ErrUserAborted {
			fmt.Println("User aborted")
			os.Exit(0)
		}
		// If there was an error, exit with an error message.
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	// Run the function to generate the credentials. With Spinner
	// _ = spinner.New().
	// 	Title("Generating Credentials...").
	// 	Action(credentials.createCredentials).Run()

	createCredentials(&credentials, output)

	var sb strings.Builder

	fmt.Fprintf(&sb, "Root: %t\n", (credentials.Root))
	fmt.Fprintf(&sb, "Username: %s\n", styles.PaintRed(credentials.Username))
	fmt.Fprintf(&sb, "Database: %s\n", styles.PaintCyan(credentials.Database))
	fmt.Fprintf(&sb, "Host: %s\n", styles.PaintBlue(credentials.Host))
	fmt.Fprintf(&sb, "Replica Set: %v\n", styles.PaintYellow(credentials.ReplicaSet))
	fmt.Fprintf(&sb, "Password: %s\n", styles.PaintGray(credentials.Password))
	fmt.Fprintf(&sb, "Authentication DB: %s\n", styles.PaintOrange(credentials.AuthenticationDB))

	fmt.Println(
		lipgloss.NewStyle().
			Width(styles.TerminalWidth-2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Render(sb.String()),
	)
}

// createCredentials generates MongoDB credentials and saves them to files.
// It performs the following steps:
// 1. Checks if a "secrets" directory exists in the same directory as the executable. If not, it creates the directory.
// 2. Logs the start of credential generation.
// 3. If the user is a ROOT user, sets the username to "root" and the database to "admin".
// 4. Converts all string fields of the MongoCredentials struct to lowercase.
// 5. Creates a MongoDB user password file with the pattern "MONGO_<USERNAME>_PW" in the specified output directory.
// 6. Checks if the password file already exists. If it does, prompts the user to confirm overwriting the file.
// 7. If the user confirms, generates a new password and writes it to the file. Otherwise, reads the existing password from the file.
// 8. Detects if the MongoDB host is part of a replica set and logs the result.
// 9. Compiles the MongoDB connection string.
// 10. Writes a summary of the credentials to a YAML file named "mongo_<username>_credentials.yaml" in the specified output directory.
//
// Parameters:
// - c: A pointer to a MongoCredentials struct containing the MongoDB credentials to be generated.
// - output: A string specifying the output directory where the credentials files will be saved.
func createCredentials(c *MongoCredentials, output string) {

	// check if exsist the output folder. If not, create it
	if _, err := os.Stat(output); os.IsNotExist(err) {
		log.Info("Creating output directory")
		if err := os.Mkdir(output, 0755); err != nil {
			log.Error(err)
			os.Exit(1)
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
				log.Error(err)
				os.Exit(1)
			}
		}
	} else {
		passwordGeneratorLogger(c, mongoPasswordFilename)
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
		log.Error(err)
		os.Exit(1)
	} else {
		log.Info(fmt.Sprintf("Credentials written %d bytes to %s", yamlBytes, credentialFilename))
	}
}

// log the password generation process while writing the password to a file
func passwordGeneratorLogger(c *MongoCredentials, mongoPasswordFilename string) {
	log.Info("Generating password...")
	if pwdBytes, err := c.createPasswordFile(mongoPasswordFilename); err != nil {
		log.Error(err)
		os.Exit(1)
	} else {
		log.Info(fmt.Sprintf("Wrote %d bytes to file %v", pwdBytes, mongoPasswordFilename))
	}
}
