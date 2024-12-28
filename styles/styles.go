package styles

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var termWidth, _, _  = term.GetSize(int(os.Stdout.Fd()))

var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#675499")).
	Padding(1).
	Margin(1, 0).
	Width(termWidth).
	Align(lipgloss.Center)

	

var CommandStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a4a8b0")).
	Margin(1, 0).
	MarginTop(0).
	Width(50).
	Align(lipgloss.Left)
	