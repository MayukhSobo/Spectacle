package savedconns

import (
	"spectacle/internal/app/common"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
)

type (
	// Connection represents a saved database connection configuration
	// with name, type, URI, and selection state. Used for both display
	// and connection management.
	Connection struct {
		Name     string
		Type     string // e.g., "etcd", "redis", "postgres"
		URI      string
		Selected bool
	}

	// SavedConnModel manages the UI state for interacting with saved connections.
	// Handles list rendering, selection, and navigation between connections.
	SavedConnModel struct {
		connections []Connection
		list        list.Model
		Window      *common.Window
		keys        *KeyMap
		help        help.Model
	}

	connectionItem Connection
)

// NewSavedConnModel creates a new saved connections model with default values.
// Initializes empty connection list and sets up UI components.
func NewSavedConnModel() SavedConnModel {
	keys := newKeyMap()
	delegate := list.NewDefaultDelegate()
	list := list.New([]list.Item{}, delegate, 0, 0)
	list.SetShowHelp(false)  // We'll use our own help
	list.SetShowTitle(false) // We'll handle the title separately

	return SavedConnModel{
		connections: make([]Connection, 0),
		list:        list,
		Window:      common.NewWindow(0, 0),
		keys:        keys,
		help:        help.New(),
	}
}

// AddConnection adds a new connection to the list
func (m *SavedConnModel) AddConnection(conn Connection) {
	m.connections = append(m.connections, conn)
	// Update the list items
	m.updateListItems()
}

// updateListItems updates the list items from connections
func (m *SavedConnModel) updateListItems() {
	items := make([]list.Item, len(m.connections))
	for i, conn := range m.connections {
		items[i] = connectionItem(conn)
	}
	m.list.SetItems(items)
}

// Implement list.Item interface for Connection
func (i connectionItem) Title() string       { return i.Name }
func (i connectionItem) Description() string { return i.URI }
func (i connectionItem) FilterValue() string { return i.Name }
