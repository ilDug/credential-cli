package styles

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var TerminalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))

var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#675499")).
	Padding(1).
	Margin(1, 0).
	MarginBottom(3).
	Width(TerminalWidth).
	Align(lipgloss.Center)

var CommandStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a4a8b0")).
	Margin(1, 0).
	MarginTop(0).
	Width(50).
	Align(lipgloss.Left)

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

