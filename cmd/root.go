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
}

var rootCmd = &cobra.Command{
	Use:   "spectacle",
	Short: "Spectacle is an ETCD explorer for your terminal",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log.Info("Starting Spectacle")
		p := tea.NewProgram(NewApp())
		if _, err := p.Run(); err != nil {
			logger.Log.Fatal(err)
		}
	},
}

func NewApp() AppModel {
	return AppModel{
		currentPage: home.NewHomeScreenModel("Welcome to Spectacle"),
	}
}

func (m AppModel) Init() tea.Cmd {
	return m.currentPage.Init()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case home.NavigateToSavedConns:
		m.currentPage = savedconns.NewSavedConnModel()
		return m, m.currentPage.Init()
	case savedconns.NavigateToHome:
		m.currentPage = home.NewHomeScreenModel("Welcome back to Spectacle")
		return m, m.currentPage.Init()
	default:
		m.currentPage, cmd = m.currentPage.Update(msg)
	}

	return m, cmd
}

func (m AppModel) View() string {
	return m.currentPage.View()
}
func Execute() error {
	return rootCmd.Execute()
}
