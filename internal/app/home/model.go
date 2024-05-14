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
}

type Tooltip struct {
	Msg    string
	Active bool
	Alert  AlertType
	// Add tooltip style
}

type Banner struct {
	BannerText         string
	BannerStatingColor string
	BannerEndingColor  string
	BannerStyle        lipgloss.Style
}

func newBanner(startCol, endCol string) *Banner {
	return &Banner{
		BannerText:         banner,
		BannerStatingColor: startCol,
		BannerEndingColor:  endCol,
	}

}

type HomeScreenModel struct {
	Banner  *Banner
	Tooltip *Tooltip
	Input   *Input
	Window  *Window
	Keys    *keyMap
	Help    help.Model
}

func NewHomeScreenModel(defaultMsg string) HomeScreenModel {
	return HomeScreenModel{
		Banner:  newBanner("#B14FFF", "#00FFA3"),
		Input:   newInput(defaultMsg),
		Window:  newWindow(0, 0), // This doesn't mean window size is 0, 0
		Keys:    newKeyMap(),
		Help:    help.New(),
		Tooltip: newTooltip(tootTipMsgGoodConn),
	}
}

func newWindow(height, width int) *Window {
	return &Window{
		Height: height,
		Width:  width,
	}
}

func newTooltip(msg string) *Tooltip {
	return &Tooltip{
		Msg:    msg,
		Active: false,
		Alert:  GOOD_CONNECTION,
	}
}

func newInput(defaultText string) *Input {
	input := new(Input)
	input.model = textinput.New()
	input.model.Placeholder = defaultText
	input.model.Focus()
	return input
}
