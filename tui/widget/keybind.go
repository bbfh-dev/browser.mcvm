package widget

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
