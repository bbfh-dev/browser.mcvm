package screen

import (
	"strings"

	"github.com/bbfh-dev/browser.mcvm/cli"
	"github.com/bbfh-dev/browser.mcvm/pkg/mcvm"
	"github.com/bbfh-dev/browser.mcvm/tui/style"
	"github.com/bbfh-dev/browser.mcvm/tui/widget"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var KEYBINDS = map[string]widget.Keybind{
	"item.next":     widget.NewKeybind("ctrl+down", "j"),
	"item.previous": widget.NewKeybind("ctrl+up", "k"),
	"item.select":   widget.NewKeybind("enter"),
}

type HomeScreen struct {
	Packages  []widget.PackageWidget
	current   int
	ready     bool
	loaded    bool
	searchFor string
}

func NewHomeScreen() HomeScreen {
	return HomeScreen{
		Packages:  []widget.PackageWidget{},
		current:   0,
		ready:     false,
		loaded:    false,
		searchFor: "",
	}
}

type loadPackagesMsg struct {
	Packages []widget.PackageWidget
}

func (screen HomeScreen) loadPackages() tea.Msg {
	packageIds := mcvm.ListAllPackages()
	var packages []widget.PackageWidget

	for _, id := range packageIds {
		pkg := widget.NewPackageWidget(id)
		pkg.Load()
		packages = append(packages, pkg)
	}

	return loadPackagesMsg{packages}
}

func (screen HomeScreen) limitCurrent() HomeScreen {
	if screen.current < 0 {
		screen.current = 0
	}

	if screen.current >= len(screen.Packages) {
		screen.current = len(screen.Packages) - 1
	}

	return screen
}

func (screen HomeScreen) Update(raw tea.Msg) (Screen, tea.Cmd) {
	if !screen.ready {
		screen.ready = true
		return screen, screen.loadPackages
	}

	var commands []tea.Cmd

	switch msg := raw.(type) {
	case loadPackagesMsg:
		screen.Packages = msg.Packages
		screen.loaded = true
	case tea.KeyMsg:
		switch {
		case KEYBINDS["item.previous"].Matches(msg):
			screen.current -= 1
		case KEYBINDS["item.next"].Matches(msg):
			screen.current += 1
		case KEYBINDS["item.select"].Matches(msg):
			// FIXME: Add actual functionality
			cli.Log(screen.Packages[screen.current].Id)
			return screen, tea.Quit
		}
	}

	screen = screen.limitCurrent()

	return screen, tea.Batch(commands...)
}

func (screen HomeScreen) View(width int) string {
	if !screen.loaded {
		return style.DefaultStyle.Width(width).
			AlignHorizontal(lipgloss.Center).
			Render(style.WithIcon(style.DATABASE_ICON, "Loading packages info..."))
	}

	if len(screen.Packages) == 0 {
		return style.DefaultStyle.Width(width).
			AlignHorizontal(lipgloss.Center).
			Render(style.WithIcon(style.DATABASE_ICON, "No data"))
	}

	var test strings.Builder

	for i, pkg := range screen.Packages {
		test.WriteString(
			style.PackageStyle(pkg, i == screen.current).RenderLine(width, pkg.Id),
		)
	}

	return test.String()
}

func (screen HomeScreen) GotoTop() Screen {
	screen.current = 0
	return screen
}

func (screen HomeScreen) GotoBottom() Screen {
	screen.current = len(screen.Packages) - 1
	return screen
}

func (screen HomeScreen) SetSearch(input string) Screen {
	screen.searchFor = input
	return screen
}
