package home

import (
	"github.com/charmbracelet/lipgloss"
)

func (i *Input) MakeStyle(w *Window) {
	centreStyle := lipgloss.NewStyle().Width(w.Width).Align(lipgloss.Center)
	padding := w.Width / 10 // This creates padding of 10% of the total width on each side
	borderWidth := w.Width - 2*padding
	borderStyle := lipgloss.NewStyle().BorderForeground(lipgloss.Color("#81b2b5")).
		BorderStyle(lipgloss.RoundedBorder()).Padding(1).Width(borderWidth)
	i.inputStyle = centreStyle
	i.borderStyle = borderStyle
	i.borderWidth = borderWidth
}

func (b *Banner) MakeStyle(w *Window) {
	centreStyle := lipgloss.NewStyle().Width(w.Width).Align(lipgloss.Center)
	b.BannerStyle = centreStyle.PaddingTop(1)
	b.RenderedBanner = GradientBanner(b)
}

func (h *Help) MakeStyle(w *Window) {
	h.style = lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(w.Width)
}

func (t *Tooltip) MakeStyle(w *Window, in *Input) {
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

func EditModeStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")). // Black text
		Background(lipgloss.Color("#FFFFFF")). // White background
		Bold(true).
		Align(lipgloss.Center).
		Width(20).    // Set a fixed width for the mode indicator
		Padding(0, 2) // Add horizontal padding to enhance the pill shape
}

func CommandModeStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")). // White text
		Background(lipgloss.Color("#457B9D")). // Darker blue background
		Bold(true).
		Align(lipgloss.Center).
		Width(20).    // Set a fixed width for the mode indicator
		Padding(0, 2) // Add horizontal padding to enhance the pill shape
}

func CenteredStyle(windowWidth int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(windowWidth).
		Align(lipgloss.Center)
}

func SpacerStyle() lipgloss.Style {
	return lipgloss.NewStyle().Height(1)
}
