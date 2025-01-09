package core

import (
	"cre/styles"

	"github.com/charmbracelet/huh"
)

var AvailableCommands = []string{"HELP", "MONGO_GENERATE", "MONGO_CREDENTIALS", "CERTIFICATE"}

type Command string

const (
	HelpCmd            Command = "HELP"
	MongoDBGenerate    Command = "MONGO_GENERATE"
	MongoDBInquire        Command = "MONGO_CREDENTIALS"
	CertificateManager Command = "CERTIFICATE"
)

func SelectCmdForm(cmd *Command) *huh.Form {

	form := huh.NewForm(
		huh.NewGroup(huh.NewSelect[Command]().
			Title("Select Command").
			Description("Select the command to run").
			Options(
				huh.NewOption("Help", HelpCmd),
				huh.NewOption("MongoDB Generate Credentials", MongoDBGenerate),
				huh.NewOption("MongoDB Inquire Credentials", MongoDBInquire),
				huh.NewOption("Certificate Manager", CertificateManager),
			).
			Value(cmd).
			// Key("cmd").
			Height(20),
		),
	).
		WithTheme(styles.ThemeDag()).
		WithHeight(20)

	return form
}
