package tui

import tea "github.com/charmbracelet/bubbletea"

type Keybind []string

func NewKeybind(binds ...string) Keybind {
	return Keybind(binds)
}

func (keybinds Keybind) Matches(msg tea.KeyMsg) bool {
	for _, bind := range keybinds {
		if msg.String() == bind {
			return true
		}
	}
	return false
}

var KEYBINDS = map[string]Keybind{
	"quit":        NewKeybind("q", "ctrl+c"),
	"goto.top":    NewKeybind("G"),
	"goto.bottom": NewKeybind("g"),
	"scroll.up":   NewKeybind("down", "J"),
	"scroll.down": NewKeybind("up", "K"),
}
