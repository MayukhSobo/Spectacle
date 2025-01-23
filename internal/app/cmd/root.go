package cmd

import (
	"spectacle/internal/app/home"
	"spectacle/internal/app/savedconns"

	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	currentPage tea.Model
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Handle window size for both pages
		if homeModel, ok := m.currentPage.(*home.HomeScreenModel); ok {
			homeModel.Window.Height = msg.Height
			homeModel.Window.Width = msg.Width
			homeModel.InitializeStyles()
		}
		if savedModel, ok := m.currentPage.(*savedconns.SavedConnModel); ok {
			savedModel.Window.Height = msg.Height
			savedModel.Window.Width = msg.Width
		}
		return m, nil

	case home.NavigateToSavedConns:
		m.currentPage = savedconns.NewSavedConnModel()
		return m, m.currentPage.Init()

	case savedconns.NavigateToHome:
		homeModel := home.NewHomeScreenModel("Welcome back to Spectacle")
		homeModel.Window.Height = m.currentPage.(*savedconns.SavedConnModel).Window.Height
		homeModel.Window.Width = m.currentPage.(*savedconns.SavedConnModel).Window.Width
		homeModel.InitializeStyles()
		m.currentPage = homeModel
		return m, m.currentPage.Init()

	default:
		m.currentPage, cmd = m.currentPage.Update(msg)
	}

	return m, cmd
}

func (m AppModel) Init() tea.Cmd {
	return m.currentPage.Init()
}

func (m AppModel) View() string {
	return m.currentPage.View()
}
