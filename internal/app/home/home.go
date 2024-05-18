package home

import (
	"math/rand"
	"os"
	"spectacle/db"
	"spectacle/logger"

	tea "github.com/charmbracelet/bubbletea"
)

func (m HomeScreenModel) Init() tea.Cmd {
	return nil
}

func ping(endpoint string, db *db.Database) (AlertType, error) {
	_ = endpoint
	_ = db
	values := []AlertType{goodConnection, noConnection}
	return values[rand.Intn(2)], nil
}

func Start() {
	m := NewHomeScreenModel("What is the ETCD endpoint?")
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		logger.Log.Errorf("Failed to start program: %v", err)
		os.Exit(1)
	}
}
