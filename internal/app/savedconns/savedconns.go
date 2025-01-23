package savedconns

import (
	"os"
	"spectacle/logger"

	tea "github.com/charmbracelet/bubbletea"
)

// Add this method to implement tea.Model interface
func (m SavedConnModel) Init() tea.Cmd {
	return nil
}

// Start initializes and runs the saved connections page
func Start() {
	m := NewSavedConnModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		logger.Log.Errorf("Failed to start saved connections page: %v", err)
		os.Exit(1)
	}
}
