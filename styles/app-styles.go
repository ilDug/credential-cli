package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)


var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).BorderBottom(true).BorderForeground(lipgloss.Color("241")).BorderStyle(lipgloss.DoubleBorder()).Foreground(lipgloss.Color("30"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).Foreground(lipgloss.Color("241")).MarginTop(0)
	quitTextStyle     = lipgloss.NewStyle().MarginLeft(4).Foreground(lipgloss.Color("110"))
)