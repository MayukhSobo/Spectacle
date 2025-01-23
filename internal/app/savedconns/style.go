package savedconns

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			PaddingLeft(2).
			PaddingBottom(1)

	listStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(0, 1)

	noConnectionsStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FAFAFA")).
				Border(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.Color("#874BFD")).
				Padding(1, 2).
				Align(lipgloss.Center)
)
