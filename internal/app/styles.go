package app

import "github.com/charmbracelet/lipgloss"

var (
	homeScreenStyle *HomeScreenStyle
)

type BannerStyleProperties struct {
	BannerGradientStartColor string
	BannerGradientEndColor   string
}

type InputFieldStyle struct {
	BorderStyle lipgloss.Border
	BorderColor lipgloss.Color
	Padding     int
	Width       int // Make it 80% of the window size
}

type HomeScreenStyle struct {
	InpFieldStyle lipgloss.Style
	BannerStyle   []lipgloss.Style
	HelpStyle     []lipgloss.Style
}

func newHomeScreenStyle(bannerProps *BannerStyleProperties) *HomeScreenStyle {
	return &HomeScreenStyle{
		InpFieldStyle: lipgloss.NewStyle(),
		BannerStyle:   makeRampStyles(bannerProps),
		HelpStyle:     nil,
	}
}

func (hss *HomeScreenStyle) AddInputFieldStyle(ifs *InputFieldStyle) {
	hss.InpFieldStyle = hss.InpFieldStyle.
		BorderForeground(ifs.BorderColor).
		BorderStyle(ifs.BorderStyle).
		Padding(ifs.Padding).Width(ifs.Width)
}
