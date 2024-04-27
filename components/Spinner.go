package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sijiramakun/seapick/utils"
)

// spinnerModel represents the model for the spinner component
type spinnerModel struct {
	spinner  spinner.Model
	quitting bool
	err      error
	message  string // Message to display
}

// NewspinnerModel initializes a new spinner model
func newspinnerModel(message string) *spinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return &spinnerModel{spinner: s, message: message}
}

// Init initializes the spinner model
func (m *spinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update updates the spinner model
func (m *spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

// View returns the view for the spinner model
func (m *spinnerModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n  %s %s ...\n\n", m.spinner.View(), m.message)
	if m.quitting {
		return str + "\n"
	}
	return str
}

func Spinner(msg string) {
	spinnerModel := newspinnerModel(msg)
	p := tea.NewProgram(spinnerModel)
	if _, err := p.Run(); err != nil {
		utils.CheckError(err)
	}
}
