package home

import (
	"spectacle/logger"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m HomeScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Window.Height = msg.Height
		m.Window.Width = msg.Width
		cmd = tea.ClearScreen
		logger.Log.Debugf("Height: %d, Width: %d", msg.Height, msg.Width)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			logger.Log.Debug("Received quit key to exit the application")
			cmd = tea.Quit
		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll
			cmd = tea.ClearScreen
			// case key.Matches(msg, m.Keys.Connect):
			// 	m.ShowTooltip = !m.ShowTooltip
			// 	return m, tea.ClearScreen
		}
	}
	return m, cmd
}
