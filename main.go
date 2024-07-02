package main

import (
	"fmt"
	"os"

	"github.com/bbfh-dev/browser.mcvm/cli"
	"github.com/bbfh-dev/browser.mcvm/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func init() {
	cli.ParseFlags()
}

func main() {
	defer cli.SetupLogger().Close()

	p := tea.NewProgram(tui.NewIndexModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
