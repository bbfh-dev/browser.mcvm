package screen

import (
	"github.com/bbfh-dev/browser.mcvm/tui/widget"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen interface {
	Update(tea.Msg) (Screen, tea.Cmd)
	View(width int) (string, []widget.Tab)
	GotoTop() Screen
	GotoBottom() Screen
	SetSearch(input string) Screen
	CurrentTab() int
	SwitchTab(int) Screen
}

const (
	HOME_SCREEN = iota
)
