package screen

import (
	"math"
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
	tab       int
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

type filterPackageFunc func(widget.PackageWidget, string) bool

func filterPackageByContents(pkg widget.PackageWidget, searchString string) bool {
	return strings.Contains(pkg.Id, searchString) ||
		strings.Contains(strings.ToLower(pkg.Name()), searchString)
}

func filterPackageByAuthor(pkg widget.PackageWidget, searchString string) bool {
	for _, author := range pkg.Authors() {
		if strings.Contains(strings.ToLower(author), searchString[1:]) {
			return true
		}
	}

	return false
}

func (screen HomeScreen) listPackages() []widget.PackageWidget {
	if widget.HasInput(screen.searchFor) {
		searchString := strings.ToLower(screen.searchFor)
		var pkgs []widget.PackageWidget

		var filterFunc filterPackageFunc

		if strings.HasPrefix(searchString, "@") {
			filterFunc = filterPackageByAuthor
		} else {
			filterFunc = filterPackageByContents
		}

		for _, pkg := range screen.Packages {
			if filterFunc(pkg, searchString) {
				pkgs = append(pkgs, pkg)
			}
		}

		return pkgs
	}

	return screen.Packages
}

func (screen HomeScreen) getPageCount() int {
	return int(math.Ceil(float64(len(screen.Packages)) / float64(widget.ITEMS_PER_PAGE)))
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

func (screen HomeScreen) View(width int) (string, []widget.Tab) {
	if !screen.loaded {
		return style.DefaultStyle.Width(width).
			AlignHorizontal(lipgloss.Center).
			Render(style.WithIcon(style.DATABASE_ICON, "Loading packages info...")), []widget.Tab{}
	}

	if len(screen.Packages) == 0 {
		return style.DefaultStyle.Width(width).
			AlignHorizontal(lipgloss.Center).
			Render(style.WithIcon(style.DATABASE_ICON, "No data")), []widget.Tab{}
	}

	var test strings.Builder

	packages := screen.listPackages()
	begin := screen.CurrentTab() * widget.ITEMS_PER_PAGE
	end := screen.CurrentTab()*widget.ITEMS_PER_PAGE + widget.ITEMS_PER_PAGE
	if end >= len(packages) {
		end = len(packages) - 1
	}

	for i, pkg := range packages[begin:end] {
		test.WriteString(
			style.PackageStyle(pkg, i == screen.current).RenderLine(width, pkg.Id),
		)
	}

	count := screen.getPageCount()
	var tabs = make([]widget.Tab, count)
	for i := range screen.getPageCount() {
		tabs[i] = widget.NewSimpleTab(i)
	}

	return test.String(), tabs
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

func (screen HomeScreen) CurrentTab() int {
	return screen.tab
}

func (screen HomeScreen) SwitchTab(tab int) Screen {
	screen.tab = tab

	if screen.tab < 0 {
		screen.tab = 0
	}

	if screen.tab >= screen.getPageCount() {
		screen.tab = screen.getPageCount() - 1
	}

	return screen
}
