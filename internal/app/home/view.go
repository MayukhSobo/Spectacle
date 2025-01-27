package home

import (
	"spectacle/internal/app/common"
	"spectacle/logger"

	"github.com/charmbracelet/lipgloss"
)

func getStyledBanner(m ScreenModel) string {
	logger.Log.Debugf("getStyledBanner is called with window width: %d\n", m.Window.Width)
	bannerModel := m.Banner
	bannerModel.MakeStyle(m.Window)
	return bannerModel.BannerStyle.Render(
		bannerModel.RenderedBanner,
	)
}

func getStyledInput(m ScreenModel) string {
	inputModel := m.Input
	inputModel.MakeStyle(m.Window)
	return inputModel.inputStyle.Render(inputModel.borderStyle.Render(
		m.Input.model.View(),
	))
}

func getStyledHelp(m ScreenModel) string {
	helpModel := m.Help
	helpModel.MakeStyle(m.Window)

	if m.Help.IsActive {
		// Show full help in command mode
		return helpModel.style.Render(m.Help.model.View(m.Keys))
	}

	// Show only toggle mode help in edit mode
	return helpModel.style.Render(m.Help.model.View(m.Keys))
}

func getStyledTooltip(m ScreenModel) string {
	tooltipModel := m.Tooltip
	inputModel := m.Input
	tooltipModel.MakeStyle(m.Window, inputModel)
	if tooltipModel.Active {
		tooltip := tooltipModel.BackgroundStyle.Render(tooltipModel.Msg)
		return tooltipModel.TooltipStyle.Render(tooltip)
	}
	return ""
}

// Spacer creates vertical spacing between UI components.
// Returns a string with the specified number of newlines.
func Spacer(w *common.Window, banner, input, tooltip, help string) string {
	totalHeight := lipgloss.Height(banner) +
		lipgloss.Height(input) +
		lipgloss.Height(tooltip)
	gapHeight := w.Height - totalHeight - lipgloss.Height(help)
	return lipgloss.NewStyle().Height(gapHeight - 1).Render("")
}

func getModeIndicator(m ScreenModel) string {
	var mode string
	var modeStyle lipgloss.Style

	if m.Help.IsActive {
		mode = "Command Mode"
		modeStyle = CommandModeStyle()
	} else {
		mode = "Edit Mode"
		modeStyle = EditModeStyle()
	}

	// Center the mode indicator within the full window width
	return CenteredStyle(m.Window.Width).Render(modeStyle.Render(mode))
}

// View renders the complete home screen UI.
// Combines banner, input field, tooltip, and help components.
// Part of tea.Model interface implementation.
func (m ScreenModel) View() string {
	// Get the banner model and apply its style
	banner := getStyledBanner(m)
	// Get the tooltip model and apply its style
	tooltip := getStyledTooltip(m)
	// Get the input model and apply its style
	input := getStyledInput(m)
	// Get the help model and apply its style
	help := getStyledHelp(m)
	// Get the mode indicator
	modeIndicator := getModeIndicator(m)

	// Create a spacer for gap using the style from style.go
	spacer := SpacerStyle().Render("")

	return lipgloss.JoinVertical(
		lipgloss.Top,
		banner,
		tooltip,
		input,
		Spacer(m.Window, banner, input, tooltip, help),
		modeIndicator, // Center the mode indicator
		spacer,        // Add a spacer for gap
		help,
	)
}
