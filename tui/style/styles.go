package style

import "github.com/charmbracelet/lipgloss"

var DefaultStyle = lipgloss.NewStyle()
var TextStyle = defaultStyle{lipgloss.NewStyle().Foreground(REGULAR_COLOR)}

var HintStyle = defaultStyle{
	lipgloss.NewStyle().Foreground(HINT_COLOR),
}

var InactiveStyle = defaultStyle{
	lipgloss.NewStyle().Foreground(INACTIVE_COLOR),
}

var HeaderStyle = defaultStyle{
	lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, true),
}

var FooterStyle = defaultStyle{
	lipgloss.NewStyle().
		Foreground(REGULAR_COLOR).
		Background(ACCENT_COLOR).
		Padding(0, 1).
		MaxHeight(1),
}

var ScrollBackground = InactiveStyle.Render(1, "┃")
var ScrollForeground = TextStyle.Render(1, "┃")
