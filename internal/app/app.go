package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

type model struct {
	question      string
	width         int
	height        int
	answerField   textinput.Model
	styles        *Styles
	choices       []string // Add this line
	choiceIndex   int      // Add this line
	displayChoice bool
}

type Styles struct {
	BorderColor    lipgloss.Color
	InputField     lipgloss.Style
	Banner         []lipgloss.Style
	Choice         lipgloss.Style // Style for unselected choices
	SelectedChoice lipgloss.Style // Style for the selected choice
}

// Helper function for converting colors to hex. Assumes a value between 0 and 1.
func colorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}

// Convert a colorful.Color to a hexadecimal format.
func colorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", colorFloatToHex(c.R), colorFloatToHex(c.G), colorFloatToHex(c.B))
}
func makeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(colorToHex(c))))
	}
	return
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("#81b2b5")
	s.InputField = lipgloss.NewStyle().
		BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).
		Width(80)
	banner := Banner()
	s.Banner = makeRampStyles("#B14FFF", "#00FFA3", float64(len(banner)))

	// Style for the selected choice
	s.SelectedChoice = lipgloss.NewStyle().
		// PaddingLeft(2).
		Foreground(lipgloss.Color("#000000")). // Example color
		Background(lipgloss.Color("#81b2b5")). // Example color
		Bold(true).PaddingRight(2).PaddingLeft(2)
	return s
}

func NewModel(question string) *model {
	styles := DefaultStyles()
	ansField := textinput.New()
	ansField.Placeholder = "your endpoint..."
	ansField.Focus()
	return &model{
		question:      question,
		answerField:   ansField,
		styles:        styles,
		choices:       []string{"Connect", "Save", "Cancel"}, // Example choices
		choiceIndex:   0,                                     // Default to the first choice
		displayChoice: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k": // Use 'k' for Vim-like keybindings
			if m.choiceIndex > 0 {
				m.choiceIndex--
			}
		case "down", "j": // Use 'j' for Vim-like keybindings
			if m.choiceIndex < len(m.choices)-1 {
				m.choiceIndex++
			}
		case "enter":
			if !m.displayChoice && m.answerField.Value() != "" {
				m.displayChoice = true
			} else if m.answerField.Value() != "" {
				if m.choiceIndex == 2 {
					m.displayChoice = false
					m.answerField.SetValue("")
				}
				m.displayChoice = false
			}
			return m, nil
		case "esc":
			m.displayChoice = false
			return m, nil
		}
	}
	m.answerField, cmd = m.answerField.Update(msg)
	return m, cmd
}
func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var bannerRendered string
	banner := Banner()
	for i, each := range m.styles.Banner {
		bannerRendered += each.Render(string(banner[i]))
	}
	bannerRendered += "\n"

	// Build choices display
	var choicesDisplay string
	for i, choice := range m.choices {
		// Apply the appropriate style based on selection
		choiceStyle := m.styles.Choice
		if i == m.choiceIndex {
			choiceStyle = m.styles.SelectedChoice
		}
		choicesDisplay += choiceStyle.Render(choice) + "\n"
	}

	// Organize the layout
	if !m.displayChoice {
		choicesDisplay = ""
	}

	form := lipgloss.JoinVertical(
		lipgloss.Center,
		bannerRendered,
		m.question,
		m.styles.InputField.Render(m.answerField.View()),
		choicesDisplay,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		form,
	)
}

func Start() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
	m := NewModel("What is the ETCD endpoint?")
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
