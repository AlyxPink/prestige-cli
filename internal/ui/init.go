package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func initScreen() tea.Msg {
	return initMsg{}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(initScreen, tea.EnterAltScreen)
}

type initMsg struct{}

type errMsg struct {
	error
}

func (e errMsg) Error() string { return e.error.Error() }
