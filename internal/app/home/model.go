package home

import (
	"spectacle/internal/app/common"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type (
	// AlertType is an enum for the different types of alerts while connecting to a database
	AlertType int

	// Input is a struct for the input field in the home screen
	Input struct {
		model       textinput.Model
		inputStyle  lipgloss.Style
		borderStyle lipgloss.Style
		borderWidth int
	}

	// Tooltip is a struct for the tooltip in the home screen
	Tooltip struct {
		Msg             string
		Active          bool
		Alert           AlertType
		TooltipStyle    lipgloss.Style
		BackgroundStyle lipgloss.Style
		// Add tooltip style
	}

	// Banner is a struct for the banner in the home screen that displays Spectacle in a gradient style
	Banner struct {
		BannerText         string
		BannerStatingColor string
		BannerEndingColor  string
		BannerStyle        lipgloss.Style
		RenderedBanner     string
	}

	// Help is a struct for the help menu in the home screen
	Help struct {
		model    help.Model
		style    lipgloss.Style
		IsActive bool
	}

	// ScreenModel is the main model for the home screen that aggregates all the models
	ScreenModel struct {
		// UI elements
		Banner  *Banner
		Tooltip *Tooltip
		Input   *Input
		Help    *Help
		// Management elements
		Window *common.Window
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

// NewHomeScreenModel creates a new home screen model by returning a HomeScreenModel struct
func NewHomeScreenModel(defaultMsg string) ScreenModel {
	return ScreenModel{
		Banner:  newBanner("#B14FFF", "#00FFA3"),
		Input:   newInput(defaultMsg),
		Window:  common.NewWindow(0, 0), // This doesn't mean window size is 0, 0
		Keys:    newKeyMap(),
		Help:    newHelp(),
		Tooltip: newTooltip(false, goodConnection),
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
func (m *ScreenModel) InitializeStyles() {
	// Initialize all component styles
	m.Banner.MakeStyle(m.Window)
	m.Input.MakeStyle(m.Window)
	m.Help.MakeStyle(m.Window)
	m.Tooltip.MakeStyle(m.Window, m.Input)
}
