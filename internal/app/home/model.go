package home

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type AlertType int

const (
	NO_CONNECTION AlertType = iota + 1
	GOOD_CONNECTION
	PENDING_CONNECTION
)

const (
	tootTipMsgGoodConn   = "✔ Connection Successful"
	tootTipMsgFailedConn = "Connection Failed"
)

type Window struct {
	Height int
	Width  int
}

type Input struct {
	model       textinput.Model
	inputStyle  lipgloss.Style
	borderStyle lipgloss.Style
	borderWidth int
}

type Tooltip struct {
	Msg             string
	Active          bool
	Alert           AlertType
	TooltipStyle    lipgloss.Style
	BackgroundStyle lipgloss.Style
	// Add tooltip style
}

type Banner struct {
	BannerText         string
	BannerStatingColor string
	BannerEndingColor  string
	BannerStyle        lipgloss.Style
	RenderedBanner     string
}

func newBanner(startCol, endCol string) *Banner {
	return &Banner{
		BannerText:         banner,
		BannerStatingColor: startCol,
		BannerEndingColor:  endCol,
	}
}

type Help struct {
	model    help.Model
	style    lipgloss.Style
	IsActive bool
}

func newHelp() *Help {
	return &Help{
		model:    help.New(),
		IsActive: false,
	}
}

type HomeScreenModel struct {
	// UI elements
	Banner  *Banner
	Tooltip *Tooltip
	Input   *Input
	Help    *Help
	// Management elements
	Window *Window
	Keys   *keyMap
}

func NewHomeScreenModel(defaultMsg string) HomeScreenModel {
	return HomeScreenModel{
		Banner:  newBanner("#B14FFF", "#00FFA3"),
		Input:   newInput(defaultMsg),
		Window:  newWindow(0, 0), // This doesn't mean window size is 0, 0
		Keys:    newKeyMap(),
		Help:    newHelp(),
		Tooltip: newTooltip(false, GOOD_CONNECTION),
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
