package tui

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbfh-dev/browser.mcvm/tui/screen"
	"github.com/bbfh-dev/browser.mcvm/tui/style"
	"github.com/bbfh-dev/browser.mcvm/tui/widget"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
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
	searchWidget widget.SearchWidget
	tabs         []widget.Tab
}

func NewIndexModel() IndexModel {
	model := IndexModel{
		Width:        MINIMUM_WIDTH,
		Height:       MINIMUM_HEIGHT,
		current:      0,
		scroll:       0,
		viewHeader:   "",
		viewContents: "",
		viewFooter:   "",
		sizeHeader:   0,
		sizeFooter:   0,
		searchWidget: widget.NewSearchWidget(),
	}

	return model
}

func (model IndexModel) renderHeader() (string, int) {
	title := " " + style.TextStyle.Render(-1, style.WithIcon(style.GAME_ICON, "MCVM Browser"))
	var rightSide string

	if !model.searchWidget.Empty() {
		rightSide = style.WithIcon(
			style.SEARCH_ICON,
			fmt.Sprintf("Search: %q", model.searchWidget.Text),
		)
	}

	version := lipgloss.PlaceHorizontal(
		model.Width-lipgloss.Width(title)-1,
		lipgloss.Right,
		style.AccentStyle.Render(-1, rightSide),
	)

	contents := style.HeaderStyle.Render(model.Width, lipgloss.JoinHorizontal(0, title, version))
	return contents, lipgloss.Height(contents)
}

func (model IndexModel) renderContents() (string, []widget.Tab) {
	return SCREENS[model.current].SetSearch(model.searchWidget.Text).View(model.Width - 1)
}

func (model IndexModel) renderFooter() (string, int) {
	var contents strings.Builder

	currentTab := SCREENS[model.current].CurrentTab()
	if len(model.tabs) > 0 {
		contents.WriteString(
			style.TabStyle.Render(
				model.Width-1,
				strings.Join(lo.Map(model.tabs, func(tab widget.Tab, index int) string {
					if index == currentTab {
						return tab.Focused()
					}
					return tab.Normal()
				}), " "),
			),
		)
		contents.WriteString("\n")
	}

	if model.searchWidget.Focused {
		contents.WriteString(style.FooterStyle.Render(
			model.Width,
			style.WithIcon(
				style.SEARCH_ICON,
				"FIND (`@` to filter by user): "+model.searchWidget.Text+"󰗧",
			),
		))
	} else {
		contents.WriteString(style.InactiveFooterStyle.Render(
			model.Width,
			style.WithIcon(style.LIST_ICON, "Carbon smashed this footer!"),
		))
	}

	result := contents.String()
	return result, lipgloss.Height(result)
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
	contents, _ := model.renderContents()
	limit := lipgloss.Height(
		contents,
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
	var eventCaptured bool

	switch msg := raw.(type) {

	case tea.WindowSizeMsg:
		model.Width, model.Height = msg.Width, msg.Height
	case tea.KeyMsg:
		// Handle as edge-case to ensure it always quits the app
		if msg.String() == "ctrl+c" {
			return model, tea.Quit
		}

		if model.searchWidget.Focused {
			model.searchWidget = model.searchWidget.HandleTyping(msg)
			eventCaptured = true
			break
		}

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
		case KEYBINDS["search"].Matches(msg):
			model.searchWidget = model.searchWidget.Focus()
		case KEYBINDS["tab.previous"].Matches(msg):
			SCREENS[model.current] = SCREENS[model.current].SwitchTab(SCREENS[model.current].CurrentTab() - 1)
		case KEYBINDS["tab.next"].Matches(msg):
			SCREENS[model.current] = SCREENS[model.current].SwitchTab(SCREENS[model.current].CurrentTab() + 1)
		}
	}
	if !eventCaptured {
		screen, cmd := SCREENS[model.current].Update(raw)
		SCREENS[model.current] = screen
		commands = append(commands, cmd)
	}

	model = model.limitScroll()

	// Render in Update() so that its able to calculate scroll properly
	header, headerHeight := model.renderHeader()
	footer, footerHeight := model.renderFooter()
	contents, tabs := model.renderContents()
	model.tabs = tabs
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
