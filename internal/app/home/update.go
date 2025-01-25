package home

import (
	"spectacle/logger"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	// ShowSavedConnsMsg signals the application to display
	// the saved connections view. Used for navigation between screens.
	ShowSavedConnsMsg struct{}

	// NavigateToSavedConns defines a message type for navigation
	NavigateToSavedConns struct{}
)

// Update handles all state changes and user interactions for the home screen.
// Processes keyboard events, manages help visibility, and handles navigation.
// Part of tea.Model interface implementation.
func (m ScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Window.Height = msg.Height
		m.Window.Width = msg.Width
		cmd = tea.ClearScreen
		logger.Log.Debugf("Height: %d, Width: %d", msg.Height, msg.Width)
	case tea.KeyMsg:
		if key.Matches(msg, m.Keys.Command) {
			logger.Log.Debugf("Toggling command mode")
			m.Help.IsActive = !m.Help.IsActive
			m.Help.model.ShowAll = m.Help.IsActive
		}
		if m.Help.IsActive {
			switch {
			case key.Matches(msg, m.Keys.Quit):
				logger.Log.Debug("Received quit key to exit the application")
				cmd = tea.Quit
			case key.Matches(msg, m.Keys.Connect):
				m.Tooltip.Active = !m.Tooltip.Active
				// Try to get the address from the input
				address := m.Input.model.Value()
				alert, err := ping(address, nil)
				if err != nil {
					logger.Log.Errorf("Encountered error: %+v", err)
					break
				}
				m.Tooltip.Alert = alert
				// Connect to the endpoint
				// If success, m.Tooltip.Alert = GOOD_CONNECTION
				// else m.Tooltip.Alert = NO_CONNECTION
				cmd = tea.ClearScreen
			case key.Matches(msg, m.Keys.Clear):
				m.Input.model.SetValue("")
				cmd = tea.ClearScreen
			case key.Matches(msg, m.Keys.Show):
				return m, func() tea.Msg {
					return NavigateToSavedConns{}
				}
			}
		} else {
			m.Input.model, cmd = m.Input.model.Update(msg)
		}
	}
	return m, cmd
}
