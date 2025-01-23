package home

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type (
	AlertType int
	Window    struct {
		Height int
		Width  int
	}
	Input struct {
		model       textinput.Model
		inputStyle  lipgloss.Style
		borderStyle lipgloss.Style
		borderWidth int
	}
	Tooltip struct {
		Msg             string
		Active          bool
		Alert           AlertType
		TooltipStyle    lipgloss.Style
		BackgroundStyle lipgloss.Style
		// Add tooltip style
	}
	Banner struct {
		BannerText         string
		BannerStatingColor string
		BannerEndingColor  string
		BannerStyle        lipgloss.Style
		RenderedBanner     string
	}
	Help struct {
		model    help.Model
		style    lipgloss.Style
		IsActive bool
	}
	HomeScreenModel struct {
		// UI elements
		Banner  *Banner
		Tooltip *Tooltip
		Input   *Input
		Help    *Help
		// Management elements
		Window *Window
		Keys   *keyMap
	}
)

const (
	noConnection AlertType = iota + 1
	goodConnection
	pendingConnection
	toolTipMsgGoodConn   = "Connection Successful :)"
	toolTipMsgFailedConn = "Connection Failed!"
)

func newBanner(startCol, endCol string) *Banner {
	return &Banner{
		BannerText:         banner,
		BannerStatingColor: startCol,
		BannerEndingColor:  endCol,
	}
}

func newHelp() *Help {
	return &Help{
		model:    help.New(),
		IsActive: false,
	}
}

func NewHomeScreenModel(defaultMsg string) HomeScreenModel {
	return HomeScreenModel{
		Banner:  newBanner("#B14FFF", "#00FFA3"),
		Input:   newInput(defaultMsg),
		Window:  newWindow(0, 0), // This doesn't mean window size is 0, 0
		Keys:    newKeyMap(),
		Help:    newHelp(),
		Tooltip: newTooltip(false, goodConnection),
	}
}

func newWindow(height, width int) *Window {
	return &Window{
		Height: height,
		Width:  width,
	}
}

func newTooltip(isActive bool, alert AlertType) *Tooltip {
	return &Tooltip{
		Active: isActive,
		Alert:  alert,
	}
}

func newInput(defaultText string) *Input {
	input := new(Input)
	input.model = textinput.New()
	input.model.Placeholder = defaultText
	input.model.Focus()
	return input
}

// InitializeStyles sets up all the styles for the home screen
func (m *HomeScreenModel) InitializeStyles() {
	// Initialize all component styles
	m.Banner.MakeStyle(m.Window)
	m.Input.MakeStyle(m.Window)
	m.Help.MakeStyle(m.Window)
	m.Tooltip.MakeStyle(m.Window, m.Input)
}
