package savedconns

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View renders the saved connections page
func (m SavedConnModel) View() string {
	var content string
	contentHeight := m.Window.Height - 6 // Reserve space for title, help, and borders

	if len(m.connections) == 0 {
		// Center the message both vertically and horizontally
		content = lipgloss.Place(
			m.Window.Width,
			contentHeight,
			lipgloss.Center,
			lipgloss.Center,
			noConnectionsStyle.Render("No saved connections yet!"),
		)
	} else {
		m.list.SetSize(
			m.Window.Width-2, // Account for borders
			contentHeight-2,  // Account for title and borders
		)
		content = fmt.Sprintf(
			"%s\n%s",
			titleStyle.Render("Saved Connections"),
			listStyle.
				Width(m.Window.Width).
				Height(contentHeight).
				Render(m.list.View()),
		)
	}

	// Add spacing and help view
	spacer := lipgloss.NewStyle().Height(1).Render("")
	return lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		spacer,
		m.help.View(m.keys),
	)
}
