package core

import (
	"cre/styles"

	"github.com/charmbracelet/huh"
)

var AvailableCommands = []string{"HELP", "MONGO_GENERATE", "MONGO_CREDENTIALS", "CERTIFICATE"}

func SelectCmdForm() *huh.Form {

	form := huh.NewForm(
		huh.NewGroup(huh.NewSelect[string]().
			Title("Select Command").
			Description("Select the command to run").
			Options(
				huh.NewOption("Help", AvailableCommands[0]),
				huh.NewOption("MongoDB Generate Credentials", AvailableCommands[1]),
				huh.NewOption("MongoDB Show Credentials", AvailableCommands[2]),
				huh.NewOption("Certificate Manager", AvailableCommands[3]),
			).
			// Value(cmd).
			Key("cmd").
			Height(14),
		),
	).
		WithTheme(styles.ThemeDag())

	return form
}
