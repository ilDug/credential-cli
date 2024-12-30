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
func MongoRun() {
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

	createCredentials(&credentials)

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

// Generate the crecentials
func createCredentials(c *MongoCredentials) {
	// check is exsta a folder called "secrets" in the same directory of the executable. If not, create it
	if _, err := os.Stat("secrets"); os.IsNotExist(err) {
		log.Info("Creating secrets directory")
		os.Mkdir("secrets", 0755)
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
	mongoPasswordFilename := fmt.Sprintf("secrets/MONGO_%s_PW", strings.ToUpper(c.Username))

	// Check if the file already exists
	if _, err := os.Stat(mongoPasswordFilename); err == nil {
		log.Warn("File already exists\n")
	}

	if overwritePW := overwriteConfim(); overwritePW {
		log.Info(("Generating password..."))
		if pwdBytes, err := c.createPasswordFile(mongoPasswordFilename); err != nil {
			log.Error(err)
			os.Exit(1)
		} else {
			log.Info(fmt.Sprintf("Wrote %d bytes to file %v", pwdBytes, mongoPasswordFilename))
		}
	} else {
		log.Warn("Skip password generation. Keeping the existing credentials")
		if err := c.readMongoPasswordFromFile(mongoPasswordFilename); err != nil {
			log.Error(err)
			os.Exit(1)
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
	credentialFilename := fmt.Sprintf("secrets/mongo_%s_credentials.yaml", strings.ToLower(c.Username))

	if yamlBytes, err := c.saveCredentialsAsYaml(credentialFilename); err != nil {
		log.Error(err)
		os.Exit(1)
	} else {
		log.Info(fmt.Sprintf("Credentials written %d bytes to %s", yamlBytes, credentialFilename))
	}
}
