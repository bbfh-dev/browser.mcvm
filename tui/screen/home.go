package screen

import tea "github.com/charmbracelet/bubbletea"

type HomeScreen struct{}

func NewHomeScreen() HomeScreen {
	return HomeScreen{}
}

func (screen HomeScreen) Update(msg tea.KeyMsg) (Screen, tea.Cmd) {
	var commands []tea.Cmd

	return screen, tea.Batch(commands...)
}

func (screen HomeScreen) View(width int, height int) string {
	return ""
}
