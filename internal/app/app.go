package app

import (
	"os"
	"spectacle/log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Use     key.Binding
	Save    key.Binding
	Connect key.Binding
	Quit    key.Binding
	Clear   key.Binding
	Show    key.Binding
	Help    key.Binding
}

// TODO: Implement input sanitization
func isSafeInput(input string) bool {
	return true
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Use, k.Save, k.Connect, k.Clear, k.Show}, // first column
		{k.Help, k.Quit}, // second column
	}
}

type WindowSize struct {
	Width  int
	Height int
}

func newKeyMap() *keyMap {
	return &keyMap{
		Use: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("return/enter", "use address"),
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
			key.WithKeys("esc"),
			key.WithHelp("esc", "quit"),
		),
		Clear: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "clear"),
		),
		Show: key.NewBinding(
			key.WithKeys("o"),
			key.WithHelp("o", "show saved addresses"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
	}
}

type InputModel struct {
	textInp      textinput.Model
	currentData  string
	currentIndex int
}

type HomeScreenModel struct {
	addressField *InputModel
	keys         *keyMap
	size         *WindowSize
	help         help.Model
}

func NewHomeScreenModel(defaultTextAddrField string) *HomeScreenModel {
	addrField := new(InputModel)
	addrField.textInp = textinput.New()
	addrField.textInp.Placeholder = defaultTextAddrField
	addrField.textInp.ShowSuggestions = true
	return &HomeScreenModel{
		addressField: addrField,
		keys:         newKeyMap(),
		size: &WindowSize{
			Width:  0,
			Height: 0,
		},
		help: help.New(),
	}
}

func DefaultInpStyle(wsize *WindowSize, frac int) *InputFieldStyle {
	return &InputFieldStyle{
		Width:       (wsize.Width * frac) / 100,
		Height: 	 (wsize.Height * frac) / 100,
		BorderColor: "#81b2b5",
		BorderStyle: lipgloss.RoundedBorder(),
		Padding:     1,
	}
}

func (m *HomeScreenModel) Init() tea.Cmd {
	homeScreenStyle = newHomeScreenStyle(&BannerStyleProperties{
		BannerGradientStartColor: "#B14FFF",
		BannerGradientEndColor:   "#00FFA3",
	})
	return nil
}

func (m *HomeScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// If the address input field is not focused,
	// focus it
	if !m.addressField.textInp.Focused() {
		m.addressField.textInp.Focus()
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Logger.Debug("Received window resize event")
		m.size.Width = msg.Width
		m.size.Height = msg.Height
		defaultInpStyle := DefaultInpStyle(m.size, 80)
		homeScreenStyle.AddInputFieldStyle(defaultInpStyle)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			log.Logger.Debug("Received quit key to exit the application")
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Clear):
			m.addressField.currentData = ""
		default:
			if msg.String() != "?" && msg.Type == tea.KeyRunes {
				m.addressField.currentData += msg.String()
			}
			if msg.Type == tea.KeyBackspace && m.addressField.currentData != "" {
				n := len(m.addressField.currentData)
				m.addressField.currentData = m.addressField.currentData[:n-1]
			}
		}

		if isSafeInput(m.addressField.currentData) {
			m.addressField.textInp.SetValue(m.addressField.currentData)
			m.addressField.textInp.CursorEnd()
		}
	}
	return m, cmd
}

func createBannerMessage(text, bgColor, textColor, emoji string) string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(textColor)). // Set text color dynamically
		Background(lipgloss.Color(bgColor)).
		PaddingTop(1).
		PaddingLeft(1).
		PaddingRight(1).
		PaddingBottom(1).
		Bold(true)
	// Apply border to all sides

	return style.Render(emoji + " " + text)
}

func (m *HomeScreenModel) View() string {
	if m.size.Width < 150 || m.size.Height < 50 {
		log.Logger.Errorf("Screen too small "+
			"for application (height = %d, width = %d)", m.size.Height, m.size.Width)
		return "Screen too small..."
	}

	m.help.View(m.keys)
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		BannerRendered(),
		homeScreenStyle.InpFieldStyle.Render(m.addressField.textInp.View()),
		// createEmptySpace((m.size.Height*50)/100),
		createBannerMessage("Connection Successful!", "#2d6a4f", "#111d13", ""),
		homeScreenStyle.HelpStyle.Render(m.help.View(m.keys)),
	)

	return lipgloss.Place(
		m.size.Width,
		m.size.Height,
		lipgloss.Center,
		0.8,
		content,
	)
}

func Start() {
	m := NewHomeScreenModel("What is the ETCD endpoint?")
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Logger.Errorf("Failed to start program: %v", err)
		os.Exit(1)
	}
}
