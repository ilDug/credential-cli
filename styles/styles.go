package styles

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var TerminalWidth, TerminalHeight, _ = term.GetSize(int(os.Stdout.Fd()))

var Base = lipgloss.NewStyle().
	Margin(1, 1, 0).
	Width(TerminalWidth - 4)

var TitleStyle = Base.
	Bold(true).
	Foreground(lipgloss.Color("#a2a2a3")).
	Background(lipgloss.Color("#676078")).
	Padding(1).
	// Margin(0).
	// MarginBottom(3).
	// Width(TerminalWidth - 2).
	Align(lipgloss.Center).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#4b4559"))

var CommandStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a4a8b0")).
	Margin(1, 0).
	MarginTop(0).
	// Width(50).
	Align(lipgloss.Left)

var BoxStyle = Base.
	Width(TerminalWidth-4).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("1")).
	Padding(1, 2)

var PaintRed = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#944753")).Render(s)
}

var PaintGreen = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#2d9574")).Render(s)
}

var PaintBlue = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#4a8cc4")).Render(s)
}

var PaintYellow = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#c4a24a")).Render(s)
}

var PaintPurple = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#675499")).Render(s)
}

var PaintGray = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#a4a8b0")).Render(s)
}

var PaintWhite = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Render(s)
}

var PaintBlack = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Render(s)
}

var PaintCyan = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#2d9574")).Render(s)
}

var PaintOrange = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#c4a24a")).Render(s)
}

var PaintMagenta = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#944753")).Render(s)
}

var PaintLightGray = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#a4a8b0")).Render(s)
}

var PaintLightWhite = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Render(s)
}

var PaintLightBlack = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Render(s)
}

var PaintLightCyan = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#2d9574")).Render(s)
}

var PaintLightOrange = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#c4a24a")).Render(s)
}

var PaintLightMagenta = func(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#944753")).Render(s)
}

type Theme struct {
	Form           lipgloss.Style
	Group          lipgloss.Style
	FieldSeparator lipgloss.Style
	Blurred        FieldStyles
	Focused        FieldStyles
	Help           help.Styles
}

// FieldStyles are the styles for input fields.
type FieldStyles struct {
	Base           lipgloss.Style
	Title          lipgloss.Style
	Description    lipgloss.Style
	ErrorIndicator lipgloss.Style
	ErrorMessage   lipgloss.Style

	// Select styles.
	SelectSelector lipgloss.Style // Selection indicator
	Option         lipgloss.Style // Select options
	NextIndicator  lipgloss.Style
	PrevIndicator  lipgloss.Style

	// FilePicker styles.
	Directory lipgloss.Style
	File      lipgloss.Style

	// Multi-select styles.
	MultiSelectSelector lipgloss.Style
	SelectedOption      lipgloss.Style
	SelectedPrefix      lipgloss.Style
	UnselectedOption    lipgloss.Style
	UnselectedPrefix    lipgloss.Style

	// Textinput and teatarea styles.
	TextInput TextInputStyles

	// Confirm styles.
	FocusedButton lipgloss.Style
	BlurredButton lipgloss.Style

	// Card styles.
	Card      lipgloss.Style
	NoteTitle lipgloss.Style
	Next      lipgloss.Style
}

// TextInputStyles are the styles for text inputs.
type TextInputStyles struct {
	Cursor      lipgloss.Style
	CursorText  lipgloss.Style
	Placeholder lipgloss.Style
	Prompt      lipgloss.Style
	Text        lipgloss.Style
}

const (
	buttonPaddingHorizontal = 2
	buttonPaddingVertical   = 0
)

// ThemeBase returns a new base theme with general styles to be inherited by
// other themes.
func ThemeBase() *huh.Theme {
	var t = &huh.Theme{}

	t.FieldSeparator = lipgloss.NewStyle().SetString("\n\n")

	button := lipgloss.NewStyle().
		Padding(buttonPaddingVertical, buttonPaddingHorizontal).
		MarginRight(1)

	// Focused styles.
	t.Focused.Base = lipgloss.NewStyle().PaddingLeft(1).BorderStyle(lipgloss.ThickBorder()).BorderLeft(true)
	t.Focused.Card = lipgloss.NewStyle().PaddingLeft(1)
	t.Focused.ErrorIndicator = lipgloss.NewStyle().SetString(" *")
	t.Focused.ErrorMessage = lipgloss.NewStyle().SetString(" *")
	t.Focused.SelectSelector = lipgloss.NewStyle().SetString("> ")
	t.Focused.NextIndicator = lipgloss.NewStyle().MarginLeft(1).SetString("→")
	t.Focused.PrevIndicator = lipgloss.NewStyle().MarginRight(1).SetString("←")
	t.Focused.MultiSelectSelector = lipgloss.NewStyle().SetString("> ")
	t.Focused.SelectedPrefix = lipgloss.NewStyle().SetString("[•] ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().SetString("[ ] ")
	t.Focused.FocusedButton = button.Foreground(lipgloss.Color("0")).Background(lipgloss.Color("7"))
	t.Focused.BlurredButton = button.Foreground(lipgloss.Color("7")).Background(lipgloss.Color("0"))
	t.Focused.TextInput.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

	t.Help = help.New().Styles

	// Blurred styles.
	t.Blurred = t.Focused
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.MultiSelectSelector = lipgloss.NewStyle().SetString("  ")
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}

func ThemeDag() *huh.Theme {
	t := ThemeBase()

	var (
		normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
		indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
		cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
		fuchsia  = lipgloss.Color("#F780E2")
		// green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
		red    = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
		orange = lipgloss.AdaptiveColor{Light: "#d19317", Dark: "#d19317"}
	)

	t.Form = t.Form.Height(25).MarginLeft(6)

	t.Focused.Base = t.Focused.Base.BorderForeground(lipgloss.Color("238")).Padding(1, 2).MarginBottom(1)
	t.Focused.Title = t.Focused.Title.Foreground(indigo).Bold(true).MarginBottom(1)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(indigo).Bold(true).MarginBottom(1)
	t.Focused.Directory = t.Focused.Directory.Foreground(indigo)
	t.Focused.Description = t.Focused.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).MarginBottom(1)
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(red)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(red)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(fuchsia)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(fuchsia)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(fuchsia)
	t.Focused.Option = t.Focused.Option.Foreground(normalFg).MarginBottom(1)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(fuchsia)
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(normalFg)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(red).Bold(true).BorderLeft(true).BorderLeftForeground(red).BorderStyle(lipgloss.NormalBorder())
	t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(normalFg)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(cream).Background(orange)
	t.Focused.Next = t.Focused.FocusedButton
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(normalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(red)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(fuchsia)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}
