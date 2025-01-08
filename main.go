package main

import (
	"cre/core"
	"cre/mongo"
	"cre/styles"
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

func main() {

	fmt.Print(styles.TitleStyle.MarginBottom(3).Render("Credentials Manager"))

	var credentials mongo.MongoCredentials
	var command core.Command

	// Run the command selector form
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

		// var out string
		// if mongoCredentialsFile != "" {
		// 	mongo.MongoCommandSelect(mongoCredentialsFile)
		// } else {
		// 	mongo.MongoRun(out)
		// }

	case core.MongoDBInquire:
		var credentialsFile string

		// load the credentials file picker form
		credentialsFilePicker := mongo.MongoFilePicker(&credentialsFile)
		err1 := credentialsFilePicker.Run()
		if err != nil {
			log.Error("Error running credentials file picker", "ERR", err1)
			os.Exit(0)
		}

		log.Info("Selected credentials file", "FILE", credentialsFile)

		// read the file contents of yaml file
		err2 := loadCredentialsFromFile(credentialsFile, &credentials)
		if err != nil {
			log.Error("Error loading credentials from file", "FILE", credentialsFile, "ERR", err2)
			os.Exit(0)
		}

		log.Info("Mongo tools ")
		mongoCmd := mongo.MongoCmd{Credentials: &credentials}
		mongo.MongoCommandSelect(&mongoCmd)

	}
}

func loadCredentialsFromFile(credentialsFile string, credentials *mongo.MongoCredentials) error {
	yamlFileContent, err := os.ReadFile(credentialsFile)
	if err != nil {
		e := errors.New("error reading file")
		e = errors.Join(e, err)
		return e
	}

	// Parse the YAML file
	err = yaml.Unmarshal(yamlFileContent, credentials)
	if err != nil {
		e := errors.New("error parsing file")
		e = errors.Join(e, err)
		return e
	}
	return nil
}
