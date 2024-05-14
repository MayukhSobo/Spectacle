package home

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Use     key.Binding
	Save    key.Binding
	Connect key.Binding
	Quit    key.Binding
	Clear   key.Binding
	Show    key.Binding
	Help    key.Binding
	Command key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Connect}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Use, k.Save, k.Connect, k.Clear, k.Show}, // first column
		{k.Help, k.Quit}, // second column
	}
}

func newKeyMap() *keyMap {
	return &keyMap{
		Use: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "use address"),
		),
		Save: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "save address"),
		),
		Connect: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "connect to endpoint"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "quit"),
		),
		Clear: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "clear"),
		),
		Show: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "list saved addresses"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Command: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "toggle command mode"),
		),
	}
}
