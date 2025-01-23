package cmd

import (
	"spectacle/internal/app/home"
	"spectacle/internal/app/savedconns"
	"spectacle/logger"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// AppModel manages the overall application state and navigation
type AppModel struct {
	currentPage tea.Model
	homeModel   home.HomeScreenModel
	savedModel  savedconns.SavedConnModel
	initialized bool
}

var rootCmd = &cobra.Command{
	Use:   "spectacle",
	Short: "Spectacle is an ETCD explorer for your terminal",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log.Info("Starting Spectacle")
		p := tea.NewProgram(NewApp(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			logger.Log.Fatal(err)
		}
	},
}

func NewApp() AppModel {
	homeModel := home.NewHomeScreenModel("Welcome to Spectacle")
	return AppModel{
		currentPage: homeModel,
		homeModel:   homeModel,
		initialized: false,
	}
}

func (m AppModel) Init() tea.Cmd {
	return m.currentPage.Init()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.initialized {
			m.initialized = true
			m.savedModel = savedconns.NewSavedConnModel()
		}
		// Forward size to current page
		m.currentPage, _ = m.currentPage.Update(msg)

	case home.NavigateToSavedConns:
		m.currentPage = m.savedModel
		return m, m.currentPage.Init()

	case savedconns.NavigateToHome:
		m.currentPage = m.homeModel
		return m, m.currentPage.Init()

	default:
		var cmd tea.Cmd
		m.currentPage, cmd = m.currentPage.Update(msg)

		// Keep the stored models in sync
		if homeModel, ok := m.currentPage.(home.HomeScreenModel); ok {
			m.homeModel = homeModel
		} else if savedModel, ok := m.currentPage.(savedconns.SavedConnModel); ok {
			m.savedModel = savedModel
		} else {
			if ok {
				logger.Log.Error("Unknown model type in currentPage")
			} else {
				logger.Log.Error("Unable to cast currentPage to any known model")
			}
		}
		return m, cmd
	}
	return m, nil
}

func (m AppModel) View() string {
	return m.currentPage.View()
}

func Execute() error {
	return rootCmd.Execute()
}
