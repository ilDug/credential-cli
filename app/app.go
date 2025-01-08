package app

import (
	"cre/core"
	"cre/mongo"
	"cre/styles"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Dag struct {
	title           lipgloss.Style
	cmdSelectorForm *huh.Form
	mongoFilePicker *huh.Form
	mongoInquire    *huh.Form
	Cmd             string
	Result          string
	credentialsFile string
	Credentials     mongo.MongoCredentials
	Tool            string
}

func (m *Dag) Init() tea.Cmd {
	// return m.cmdSelectorForm.Init()
	return tea.Batch(
		m.cmdSelectorForm.Init(),
		m.mongoFilePicker.Init(),
	)
}

////////////////////////////////////////////////////////////////////////

func (m *Dag) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// WINDOW RESIZE
	case tea.WindowSizeMsg:
		//Header
		h, _ := styles.TitleStyle.GetFrameSize()
		m.title = m.title.Width(msg.Width - h)

		return m, nil

	// KEYBOARD COMMANDS
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	// commands buffer
	var cmds []tea.Cmd

	if m.cmdSelectorForm.State == huh.StateCompleted {

		switch m.Cmd {
		case "HELP":

		case "MONGO_CREDENTIALS":

			if m.mongoFilePicker.State == huh.StateCompleted {
				// Mongo Inquire Tool
				mongoInquireForm, mongoInquireCmd := m.mongoInquire.Update(msg)
				if f, ok := mongoInquireForm.(*huh.Form); ok {
					m.mongoInquire = f
					cmds = append(cmds, mongoInquireCmd)
				}

                
			}

			// Mongo File Picker
			mongoFilePickerForm, mongoFilePickerCmd := m.mongoFilePicker.Update(msg)
			if f, ok := mongoFilePickerForm.(*huh.Form); ok {
				m.mongoFilePicker = f
				cmds = append(cmds, mongoFilePickerCmd)
			}
		}
	}

	selectorForm, selectorCmd := m.cmdSelectorForm.Update(msg)
	if f, ok := selectorForm.(*huh.Form); ok {
		m.cmdSelectorForm = f
		cmds = append(cmds, selectorCmd)
	}

	// assign the command to the model
	m.Cmd = m.cmdSelectorForm.GetString("cmd")

	return m, tea.Batch(cmds...)
}

////////////////////////////////////////////////////////////////////////

func (m *Dag) View() string {
	var sb strings.Builder

	// Header
	sb.WriteString(m.title.Render("Credential Manager"))

	// Command Selector Form
	sb.WriteString(m.cmdSelectorForm.View())

	switch m.Cmd {
	case "HELP":
		// Help
		sb.WriteString(styles.BoxStyle.Render(core.Help))

	case "MONGO_CREDENTIALS":
		// Mongo File Picker
		sb.WriteString(m.mongoFilePicker.View())

		if m.mongoFilePicker.State == huh.StateCompleted {
			// Mongo Inquire Tool
            m.mongoInquire.NextField()
			sb.WriteString(m.mongoInquire.View())
		}
	}

	return sb.String()
}

////////////////////////////////////////////////////////////////////////

func Program() *tea.Program {
	m := Dag{
		title:           styles.TitleStyle,
		credentialsFile: ".",
	}

	app := tea.NewProgram(&m, tea.WithAltScreen())

	m.cmdSelectorForm = core.SelectCmdForm()
	m.mongoFilePicker = mongo.MongoFilePicker(&m.credentialsFile)
	m.mongoInquire = mongo.MongoInquireForm(&m.Tool)

	return app
}
