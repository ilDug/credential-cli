package certificates

import (
	"cre/styles"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func customKeyMap() *huh.KeyMap {
	var dkm *huh.KeyMap = huh.NewDefaultKeyMap()
	dkm.FilePicker = huh.FilePickerKeyMap{
		GoToTop:  key.NewBinding(key.WithKeys("g"), key.WithHelp("g", "first"), key.WithDisabled()),
		GoToLast: key.NewBinding(key.WithKeys("G"), key.WithHelp("G", "last"), key.WithDisabled()),
		PageUp:   key.NewBinding(key.WithKeys("K", "pgup"), key.WithHelp("pgup", "page up"), key.WithDisabled()),
		PageDown: key.NewBinding(key.WithKeys("J", "pgdown"), key.WithHelp("pgdown", "page down"), key.WithDisabled()),
		Back:     key.NewBinding(key.WithKeys("h", "backspace", "left", "esc"), key.WithHelp("h", "back"), key.WithDisabled()),
		Select:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
		Up:       key.NewBinding(key.WithKeys("up", "k", "ctrl+k", "ctrl+p"), key.WithHelp("↑", "up"), key.WithDisabled()),
		Down:     key.NewBinding(key.WithKeys("down", "j", "ctrl+j", "ctrl+n"), key.WithHelp("↓", "down"), key.WithDisabled()),

		Open:   key.NewBinding(key.WithKeys("tab", "right"), key.WithHelp("tab", "open")),
		Close:  key.NewBinding(key.WithKeys("esc", "q", "enter"), key.WithHelp("esc, q", "close")),
		Prev:   key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back"), key.WithDisabled()),
		Next:   key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next"), key.WithDisabled()),
		Submit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
	}
	return dkm
}

func CrtPathPicker(path *string) *huh.Form {

	km := customKeyMap()

	var form = huh.NewForm(
		huh.NewGroup(
			huh.NewFilePicker().
				Picking(true).
				Title("Certificate path").
				Description("Select a directory").
				CurrentDirectory(*path).
				AllowedTypes([]string{}).
				ShowHidden(true).
				FileAllowed(false).
				DirAllowed(true).
				//
				Value(path),
		),
	).
		WithShowHelp(true).
		WithTheme(styles.ThemeDag()).
		WithKeyMap(km).
		WithHeight(20)

	return form
}
