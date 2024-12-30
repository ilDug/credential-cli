package mongo

import (
	"cre/styles"

	"github.com/charmbracelet/huh"
)

func overwriteConfim() bool {
	var overwrite bool

	f := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Overwrite").
				Description("Password file already exixts. Do you want to overwrite the existing password file?").
				Value(&overwrite).
				Affirmative("Yes").
				Negative("No"),
		),
	).WithHeight(20).WithTheme(styles.ThemeDag())
	//.WithProgramOptions(tea.WithAltScreen())

	f.Run()

	return overwrite
}
