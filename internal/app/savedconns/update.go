package savedconns

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// NavigateToHome signals a request to return to the home screen.
// Used for handling navigation between different views.
type NavigateToHome struct{}

// Update handles the update of the saved connections page
func (m SavedConnModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Window.Height = msg.Height
		m.Window.Width = msg.Width
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4) // Subtract space for title and borders

	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Back) {
			return m, func() tea.Msg {
				return NavigateToHome{}
			}
		} else if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}
	}

	var listCmd tea.Cmd
	m.list, listCmd = m.list.Update(msg)
	return m, listCmd
}
