package savedconns

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the saved connections model.
// Implements tea.Model interface for Bubble Tea initialization.
// Returns nil command as connection loading happens during model creation.
func (m SavedConnModel) Init() tea.Cmd {
	return nil
}
