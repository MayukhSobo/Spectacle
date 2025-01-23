package home

import (
	"crypto/rand"
	"math/big"
	"spectacle/db"

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
