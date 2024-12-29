package core

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// definisce lo stile della lista
var docStyle = lipgloss.NewStyle().Margin(1, 2)

// definisce la struttura degli ITEM della lista
type command struct {
	title string
    description string
    code string
}

func (i command) Title() string       { return i.title }
func (i command) Description() string { return i.description }
func (i command) FilterValue() string { return i.title }
func (i command) Code() string        { return i.code }

////////////////////////////////////////////////////////////////

// definisce il modello TEA della lista
type model struct {
	list list.Model
    selectedCommand string
}

// inizializza il modello TEA
func (m model) Init() tea.Cmd {
	return nil
}

// definisce la funzione di update del modello TEA
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
        case tea.KeyMsg:
            // if msg.String() == "ctrl+c" {
            //     return m, tea.Quit
            // }
            switch msg.String() {
                case "ctrl+c", "q":
                    return m, tea.Quit
                case "enter":
                    if m.list.SelectedItem() != nil {
                        m.selectedCommand = m.list.SelectedItem().(command).Code()
                        return m, tea.Quit
                    }
            }
        case tea.WindowSizeMsg:
            h, v := docStyle.GetFrameSize()
            m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// definisce la funzione di view del modello TEA
func (m model) View() string {
	return docStyle.Render(m.list.View())
}

////////////////////////////////////////////////////////////////

// seleziona il comando da eseguire e lo restituisce
func SelectCommand() string {
	items := []list.Item{
		command{title: "Mongo Credentials", description: "Generate a new set of credentials for MongoDB", code: "MONGO"},
		command{title: "Certificate", description: "Generate a new certificate key pair", code: "CERTIFICATE"},
	}

    // Creazione del modello TEA
	listModel := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	listModel.list.Title = "Select Command"

    // Esecuzione del programma TEA
	p := tea.NewProgram(listModel, tea.WithAltScreen())
    
    selectionModel,err := p.Run()

	if  err != nil {
		fmt.Println("Error running command selection:", err)
		os.Exit(1)
	}

    return selectionModel.(model).selectedCommand
}
