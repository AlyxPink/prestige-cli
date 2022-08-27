package utils

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	PrevLayer    key.Binding
	NextLayer    key.Binding
	PrevUpgrade  key.Binding
	NextUpgrade  key.Binding
	ManualAction key.Binding
	BuyUpgrade   key.Binding
	SellUpgrade  key.Binding
	PageDown     key.Binding
	PageUp       key.Binding
	Help         key.Binding
	Quit         key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevLayer, k.NextLayer},
		{k.PageDown, k.PageUp},
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
	PrevUpgrade: key.NewBinding(
		key.WithKeys("left", "s"),
		key.WithHelp("←/s", "move left"),
	),
	NextUpgrade: key.NewBinding(
		key.WithKeys("right", "d"),
		key.WithHelp("→/d", "move right"),
	),
	ManualAction: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↵", "enter"),
	),
	BuyUpgrade: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "Buy upgrade"),
	),
	SellUpgrade: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "Sell upgrade"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("ctrl+u"),
		key.WithHelp("Ctrl+u", "preview page up"),
	),
	PageDown: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("Ctrl+d", "preview page down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
