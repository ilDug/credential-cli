package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var TitleSyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#675499")).
	Padding(1).
	Margin(1,0).
	Width(50).
	Align(lipgloss.Center)