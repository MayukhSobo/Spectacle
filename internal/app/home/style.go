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
}

func (b *Banner) MakeStyle(w *Window) {
	centreStyle := lipgloss.NewStyle().Width(w.Width).Align(lipgloss.Center)
	b.BannerStyle = centreStyle.PaddingTop(1)
	gradientBanner := GradientBanner(
		b.BannerStatingColor,
		b.BannerEndingColor,
		b.BannerText,
	)
	b.BannerText = gradientBanner
}
