package utils

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	PrevLayer key.Binding
	NextLayer key.Binding
	Prestige  key.Binding
	Help      key.Binding
	Quit      key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevLayer, k.NextLayer},
		{k.Prestige},
		{k.Help, k.Quit},
	}
}

var Keys = KeyMap{
	PrevLayer: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	NextLayer: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Prestige: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "prestige"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
