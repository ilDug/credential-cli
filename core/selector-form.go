package core

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

// definisce la struttura degli ITEM della lista


func SelectCmd() string{
	var cmd string

	form := huh.NewForm(
		huh.NewGroup(huh.NewSelect[string]().
			Title("Select Command").
			Description("Select the command to run").
			Options(
				huh.NewOption("MongoDB Utilities", "MONGO"),
				huh.NewOption("Certificate Manager", "CERTIFICATE"),
			).
			Value(&cmd),
		),
	)

	err := form.Run()

	if err != nil {
		fmt.Println("Error running command selection:", err)
		os.Exit(1)
	}

	return cmd
}