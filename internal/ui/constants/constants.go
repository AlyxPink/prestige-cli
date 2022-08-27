package constants

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	NextView key.Binding
	PrevView key.Binding
	Help     key.Binding
	Quit     key.Binding
}

type Dimensions struct {
	Width  int
	Height int
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	PrevView: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("/h", "previous view"),
	),
	NextView: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("/l", "next view"),
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
