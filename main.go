package main

import (
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
	case "MONGO":
		fmt.Printf("Output directory: %v\n", out)
		mongo.MongoRun(out)

	case "CERTIFICATE":
		fmt.Println(styles.CommandStyle.Render("Certificate Manager... generating certificate/key pair"))
		fmt.Printf("Output directory: %v\n", out)

	}
}
