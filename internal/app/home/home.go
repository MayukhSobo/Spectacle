package home

import (
	"crypto/rand"
	"math/big"
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
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(values))))
	if err != nil {
		return noConnection, err
	}
	return values[n.Int64()], nil
}

func Start() {
	m := NewHomeScreenModel("What is the ETCD endpoint?")
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		logger.Log.Errorf("Failed to start program: %v", err)
		os.Exit(1)
	}
}
