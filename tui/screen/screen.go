package screen

import tea "github.com/charmbracelet/bubbletea"

type Screen interface {
	Update(tea.KeyMsg) (Screen, tea.Cmd)
	View(width int, height int) string
}

const (
	HOME_SCREEN = iota
)
