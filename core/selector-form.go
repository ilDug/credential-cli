package core

import (
	"cre/styles"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)


var AvailableCommands = []string{"HELP", "MONGO", "CERTIFICATE"}

func SelectCmd() string {
	var cmd string

	form := huh.NewForm(
		huh.NewGroup(huh.NewSelect[string]().
			Title("Select Command").
			Description("Select the command to run").
			Options(
				huh.NewOption("Help", AvailableCommands[0]),
				huh.NewOption("MongoDB Utilities", AvailableCommands[1]),
				huh.NewOption("Certificate Manager", AvailableCommands[2]),
			).
			Value(&cmd),
		),
	)

	err := form.WithTheme(styles.ThemeDag()).Run()

	if err != nil {
		fmt.Println("Error running command selection:", err)
		os.Exit(1)
	}

	return cmd
}
