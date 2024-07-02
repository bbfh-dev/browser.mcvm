package tui

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbfh-dev/browser.mcvm/tui/screen"
	"github.com/bbfh-dev/browser.mcvm/tui/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const MINIMUM_WIDTH = 32
const MINIMUM_HEIGHT = 10

var SCREENS = map[uint8]screen.Screen{
	screen.HOME_SCREEN: screen.NewHomeScreen(),
}

type IndexModel struct {
	Width        int
	Height       int
	current      uint8
	scroll       int
	viewHeader   string
	viewContents string
	viewFooter   string
	sizeHeader   int
	sizeFooter   int
}

func NewIndexModel() IndexModel {
	model := IndexModel{
		Width:   MINIMUM_WIDTH,
		Height:  MINIMUM_HEIGHT,
		current: 0,
	}

	return model
}

func (model IndexModel) renderHeader() (string, int) {
	title := " " + style.TextStyle.Render(-1, style.WithIcon(style.GAME_ICON, "MCVM Browser"))
	version := lipgloss.PlaceHorizontal(
		model.Width-lipgloss.Width(title)-1,
		lipgloss.Right,
		style.HintStyle.Render(-1, "by bbfh"),
	)

	contents := style.HeaderStyle.Render(model.Width, lipgloss.JoinHorizontal(0, title, version))
	return contents, lipgloss.Height(contents)
}

func (model IndexModel) renderContents() string {
	return SCREENS[model.current].View(model.Width - 1)
}

func (model IndexModel) renderFooter() (string, int) {
	contents := style.FooterStyle.Render(
		model.Width,
		style.WithIcon(style.LIST_ICON, "0 Packages installed"),
	)

	return contents, lipgloss.Height(contents)
}

func (model IndexModel) scrollStart() int {
	return max(0, model.scroll)
}

func (model IndexModel) scrollEnd(contents []string, height int) int {
	return max(0, min(len(contents), height+model.scroll))
}

func (model IndexModel) toScrollableWindow(contents string, height int) string {
	contentLines := strings.Split(strings.TrimSuffix(contents, "\n"), "\n")
	contentsHeight := len(contentLines)

	visibleContents := contentLines[model.scrollStart():model.scrollEnd(contentLines, height)]
	visibleHeight := len(visibleContents)

	barRatio := float64(height) / float64(contentsHeight)
	barHeight := int(math.Max(float64(height)*barRatio, 1))
	barPos := int(math.Round(float64(model.scroll) * barRatio))

	if barPos == 0 && model.scroll != 0 {
		barPos = 1
	}

	if model.scroll+height >= contentsHeight {
		barPos = height - barHeight
	}

	for range height - visibleHeight {
		visibleContents = append(visibleContents, strings.Repeat(" ", model.Width-1))
	}

	var builder strings.Builder

	for i, line := range visibleContents {
		line += strings.Repeat(" ", max(0, model.Width-1-len(line)))
		if i >= barPos && i < barHeight+barPos {
			builder.WriteString(line + style.ScrollForeground + "\n")
		} else {
			builder.WriteString(line + style.ScrollBackground + "\n")
		}
	}

	return builder.String()
}

func (model IndexModel) limitScroll() IndexModel {
	limit := lipgloss.Height(
		model.renderContents(),
	) - (model.Height - model.sizeHeader - model.sizeFooter)

	if model.scroll >= limit {
		model.scroll = limit
	}

	if model.scroll < 0 {
		model.scroll = 0
	}

	return model
}

func (model IndexModel) Init() tea.Cmd {
	return nil
}

func (model IndexModel) Update(raw tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd

	switch msg := raw.(type) {

	case tea.WindowSizeMsg:
		model.Width, model.Height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch {
		case KEYBINDS["quit"].Matches(msg):
			return model, tea.Quit
		case KEYBINDS["scroll.up"].Matches(msg):
			model.scroll += 1
		case KEYBINDS["scroll.down"].Matches(msg):
			model.scroll -= 1
		case KEYBINDS["goto.top"].Matches(msg):
			model.scroll = 0
			SCREENS[model.current] = SCREENS[model.current].GotoTop()
		case KEYBINDS["goto.bottom"].Matches(msg):
			model.scroll = math.MaxInt32
			SCREENS[model.current] = SCREENS[model.current].GotoBottom()
		}
	}
	screen, cmd := SCREENS[model.current].Update(raw)
	SCREENS[model.current] = screen
	commands = append(commands, cmd)

	model = model.limitScroll()

	// Render in Update() so that its able to calculate scroll properly
	header, headerHeight := model.renderHeader()
	footer, footerHeight := model.renderFooter()
	contents := model.renderContents()
	scrollableWindow := model.toScrollableWindow(contents, model.Height-headerHeight-footerHeight-1)
	model.viewHeader = header
	model.viewFooter = footer
	model.viewContents = scrollableWindow
	model.sizeHeader = headerHeight
	model.sizeFooter = footerHeight

	return model, tea.Batch(commands...)
}

func (model IndexModel) View() string {
	if model.Width < MINIMUM_WIDTH || model.Height < MINIMUM_HEIGHT {
		return fmt.Sprintf("Window is too small %v %v", model.Width, model.Height)
	}

	return lipgloss.JoinVertical(
		0,
		model.viewHeader,
		model.viewContents,
		model.viewFooter,
	)
}
