package home

import (
	"github.com/charmbracelet/lipgloss"
)

func getStyledBanner(m *HomeScreenModel) string {
	bannerModel := m.Banner
	bannerModel.MakeStyle(m.Window)
	return bannerModel.BannerStyle.Render(
		bannerModel.RenderedBanner,
	)
}

func getStyledInput(m *HomeScreenModel) string {
	inputModel := m.Input
	inputModel.MakeStyle(m.Window)
	return inputModel.inputStyle.Render(inputModel.borderStyle.Render(
		m.Input.model.View(),
	))
}

func (m HomeScreenModel) View() string {
	// Get the banner model and apply its style
	banner := getStyledBanner(&m)
	// Get the input model and apply its style
	input := getStyledInput(&m)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		banner,
		input,
	)
}
