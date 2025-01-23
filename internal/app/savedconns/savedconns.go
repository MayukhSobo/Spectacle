package savedconns

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Add this method to implement tea.Model interface
func (m SavedConnModel) Init() tea.Cmd {
	return nil
}
