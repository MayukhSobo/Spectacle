package home

import (
	"spectacle/internal/app/common"

	"github.com/charmbracelet/lipgloss"
)

// MakeStyle configures input field styling based on window dimensions.
// Applies centered alignment and creates a bordered container with padding.
func (i *Input) MakeStyle(w *common.Window) {
	centreStyle := lipgloss.NewStyle().Width(w.Width).Align(lipgloss.Center)
	padding := w.Width / 10 // This creates padding of 10% of the total width on each side
	borderWidth := w.Width - 2*padding
	borderStyle := lipgloss.NewStyle().BorderForeground(lipgloss.Color("#81b2b5")).
		BorderStyle(lipgloss.RoundedBorder()).Padding(1).Width(borderWidth)
	i.inputStyle = centreStyle
	i.borderStyle = borderStyle
	i.borderWidth = borderWidth
}

// MakeStyle configures banner styling with centered alignment and gradient.
// Updates banner style and renders the gradient banner text.
func (b *Banner) MakeStyle(w *common.Window) {
	centreStyle := lipgloss.NewStyle().Width(w.Width).Align(lipgloss.Center)
	b.BannerStyle = centreStyle.PaddingTop(1)
	b.RenderedBanner = GradientBanner(b)
}

// MakeStyle configures help panel styling with centered alignment.
// Adjusts width based on window dimensions.
func (h *Help) MakeStyle(w *common.Window) {
	h.style = lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(w.Width)
}

// MakeStyle configures tooltip appearance based on connection status.
// Adjusts colors, width, and positioning relative to input field.
func (t *Tooltip) MakeStyle(w *common.Window, in *Input) {
	borderWidth := in.borderWidth
	tooltipPadding := borderWidth / 10
	tooltipWidth := borderWidth - 7*tooltipPadding
	var backgroundColor lipgloss.AdaptiveColor
	var foregroundColor lipgloss.AdaptiveColor
	switch t.Alert {
	case goodConnection:
		backgroundColor = lipgloss.AdaptiveColor{Light: "#2E5930", Dark: "#13543C"}
		foregroundColor = lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#FFFFFF"}
		t.Msg = toolTipMsgGoodConn
	case noConnection:
		backgroundColor = lipgloss.AdaptiveColor{Light: "#7F1A1A", Dark: "#CC1B6B"}
		foregroundColor = lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#FFFFFF"}
		t.Msg = toolTipMsgFailedConn
	}
	t.BackgroundStyle = lipgloss.NewStyle().
		Background(backgroundColor).Foreground(foregroundColor).
		Bold(true).Padding(1).Align(lipgloss.Center).Width(tooltipWidth)
	t.TooltipStyle = lipgloss.NewStyle().Width(w.Width).Align(lipgloss.Center)
}

// EditModeStyle returns styling for edit mode indicator.
// Creates a white pill-shaped background with black text.
func EditModeStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")). // Black text
		Background(lipgloss.Color("#FFFFFF")). // White background
		Bold(true).
		Align(lipgloss.Center).
		Width(20).    // Set a fixed width for the mode indicator
		Padding(0, 2) // Add horizontal padding to enhance the pill shape
}

// CommandModeStyle returns styling for command mode indicator.
// Creates a blue pill-shaped background with white text.
func CommandModeStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")). // White text
		Background(lipgloss.Color("#457B9D")). // Darker blue background
		Bold(true).
		Align(lipgloss.Center).
		Width(20).    // Set a fixed width for the mode indicator
		Padding(0, 2) // Add horizontal padding to enhance the pill shape
}

// CenteredStyle creates a style that centers content horizontally.
// Takes window width to ensure proper alignment across the screen.
func CenteredStyle(windowWidth int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(windowWidth).
		Align(lipgloss.Center)
}

// SpacerStyle creates a style for vertical spacing between components.
// Returns a style with fixed height for consistent spacing.
func SpacerStyle() lipgloss.Style {
	return lipgloss.NewStyle().Height(1)
}
