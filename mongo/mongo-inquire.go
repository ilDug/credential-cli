package mongo

import (
	"cre/styles"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).BorderBottom(true).BorderForeground(lipgloss.Color("241")).BorderStyle(lipgloss.DoubleBorder()).Foreground(lipgloss.Color("30"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).Foreground(lipgloss.Color("241")).MarginTop(0)
	quitTextStyle     = lipgloss.NewStyle().MarginLeft(4).Foreground(lipgloss.Color("110"))
)

type item struct {
	name string
	fn   func() string
}

func (i item) FilterValue() string { return i.name }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("[•] %s", i.name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list   list.Model
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.fn()
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var sb strings.Builder
	sb.WriteString("\n" + m.list.View())
	sb.WriteString("\n\n" + helpStyle.Render("//////////////////////////////////////////////////////////////////\n"))
	sb.WriteString(quitTextStyle.Render(m.choice))
	sb.WriteString(helpStyle.Render("//////////////////////////////////////////////////////////////////"))

	return sb.String()
}

func MongoCommandSelect(cmd *MongoCmd) {
	items := []list.Item{
		item{"Create Mongo User", cmd.CreateUser},
		item{"Create Mongo Root User", cmd.CreateRootUser},
		item{"Mongo Connection String", cmd.ConnectionString},
		item{"Drop Mongo User", cmd.DropUser},
		item{"Authenticate Mongo User", cmd.Authenticate},
		item{"Change Mongo User Password", cmd.ChangePassword},
		item{"Grant Role to User", cmd.GrantRolesToUser},
		item{"MongoShell Cmd", cmd.MongoShellCmd},
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select the desired command"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	if res, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	} else {
		fmt.Println(styles.CommandStyle.Render(res.(model).choice))
	}

}
