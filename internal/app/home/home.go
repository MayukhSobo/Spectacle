package home

import (
	"crypto/rand"
	"math/big"
	"spectacle/db"

	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the home screen model.
// Implements tea.Model interface for Bubble Tea initialization.
// Returns nil command as no initial async operations are needed.
func (m ScreenModel) Init() tea.Cmd {
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
