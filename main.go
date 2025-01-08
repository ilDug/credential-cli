package main

import (
	"cre/core"
	"cre/styles"
	"fmt"

	"github.com/charmbracelet/log"
)

func main() {

	fmt.Print(styles.TitleStyle.MarginBottom(3).Render("Credentials Manager"))

	// var mongoCredentialsFile string
	var command core.Command
	commandSelectorForm := core.SelectCmdForm(&command)
	
	err := commandSelectorForm.Run()
	if err != nil {
		fmt.Println("Error running command selection:", err)
	}
	
	log.Info("Selected command", "CMD", command)

	switch command {
	case core.HelpCmd:
		fmt.Println(styles.BoxStyle.Render(core.Help))

	case core.MongoDBGenerate:

		// if mongoCredentialsFile != "" {
		// 	mongo.MongoCommandSelect(mongoCredentialsFile)
		// } else {
		// 	mongo.MongoRun(out)
		// }

	case core.MongoDBShow:
		// fmt.Println(styles.CommandStyle.Render("Certificate Manager... generating certificate/key pair"))

		// certPath := "secrets/certs"
		// keyPath := "secrets/keys"

		// if err := certificates.GenerateCertificate(&certPath, &keyPath); err != nil {
		// 	log.Fatal(err)
		// }

	}

	// Define command-line flags with short alternatives
	// var out string
	// flag.StringVar(&out, "out", "./secrets", "Output directory")
	// flag.StringVar(&out, "o", "./secrets", "Output directory (shorthand)")

	// flag.StringVar(&mongoCredentialsFile, "credentials", "", "file of credentials")
	// flag.StringVar(&mongoCredentialsFile, "c", "", "file of credentials (shorthand)")

	// flag.Parse()

	// Get the command argument
	// args := flag.Args()
	// if len(args) == 0 {
	// 	command = core.SelectCmd()
	// } else if len(args) > 1 {
	// 	log.Fatal("Too many arguments.", "\nSOLUTION", "Please put OPTIONS before ARGUMENTS, or be sure to provide one or no argument.")
	// } else {
	// 	command = strings.ToUpper(args[0])
	// 	if !slices.Contains(core.AvailableCommands, command) {
	// 		command = core.SelectCmd()
	// 	}
	// }

}
