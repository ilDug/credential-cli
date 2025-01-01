package main

import (
	"cre/certificates"
	"cre/core"
	"cre/mongo"
	"cre/styles"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	var command string

	// Define command-line flags with short alternatives
	var out string
	flag.StringVar(&out, "out", "./secrets", "Output directory")
	flag.StringVar(&out, "o", "./secrets", "Output directory (shorthand)")

	var mongoCredentialsFile string
	flag.StringVar(&mongoCredentialsFile, "credentials", "", "file of credentials")
	flag.StringVar(&mongoCredentialsFile, "c", "", "file of credentials (shorthand)")

	flag.Parse()

	// Get the command argument
	args := flag.Args()
	if len(args) == 0 {
		command = core.SelectCmd()
	} else if len(args) > 1 {
		log.Fatal("Too many arguments.", "\nSOLUTION", "Please put OPTIONS before ARGUMENTS, or be sure to provide one or no argument.")
	} else {
		command = strings.ToUpper(args[0])
		if !slices.Contains(core.AvailableCommands, command) {
			command = core.SelectCmd()
		}
	}

	fmt.Print(styles.TitleStyle.Render("Credentials Manager"))
	switch command {
	case "HELP":
		help := `
Utility to manage credentials for MongoDB and generate certificates.

Usage: cre [OPTIONS] COMMAND
Available commands:
help         Show this help message. 
mongo        Manage MongoDB credentials. OPTIONS: -credentials, -c | -out, -o
certificate  Generate a certificate/key pair. OPTIONS: -out, -o

Options:
-out, -o         Output directory (default: ./secrets)
-credentials, -c Path to credentials file. show commands to use in mongoShell and info about user.
		`
		fmt.Println(styles.BoxStyle.Render(help))

	case "MONGO":

		if mongoCredentialsFile != "" {
			mongo.MongoCommandSelect(mongoCredentialsFile)
		} else {
			mongo.MongoRun(out)
		}

	case "CERTIFICATE":
		fmt.Println(styles.CommandStyle.Render("Certificate Manager... generating certificate/key pair"))

		certPath := "secrets/certs"
		keyPath := "secrets/keys"

		if err := certificates.GenerateCertificate(&certPath, &keyPath); err != nil {
			log.Fatal(err)
		}

	}
}
