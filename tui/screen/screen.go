package screen

import tea "github.com/charmbracelet/bubbletea"

type Screen interface {
	Update(tea.Msg) (Screen, tea.Cmd)
	View(width int) string
	GotoTop() Screen
	GotoBottom() Screen
}

const (
	HOME_SCREEN = iota
)
