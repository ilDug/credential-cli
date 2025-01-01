package core

var Help string = `
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
