package home

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Use     key.Binding
	Save    key.Binding
	Connect key.Binding
	Quit    key.Binding
	Clear   key.Binding
	Show    key.Binding
	Command key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	// Only show toggle mode help in edit mode
	return []key.Binding{k.Command}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Use, k.Save, k.Connect},
		{k.Clear, k.Show, k.Quit},
		{k.Command},
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
			key.WithHelp("c", "connect"),
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
		Command: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "toggle mode"),
		),
	}
}
