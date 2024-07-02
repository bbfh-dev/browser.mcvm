package style

import "github.com/charmbracelet/lipgloss"

var DefaultStyle = lipgloss.NewStyle()
var TextStyle = defaultStyle{lipgloss.NewStyle().Foreground(REGULAR_COLOR)}

var HintStyle = defaultStyle{
	lipgloss.NewStyle().Foreground(HINT_COLOR),
}

var AccentStyle = defaultStyle{
	lipgloss.NewStyle().Foreground(ACCENT_COLOR),
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

func PackageStyle(card Card, selected bool) cardStyle {
	return cardStyle{
		TopStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true, true, false, true).
			Padding(0, 1).
			Foreground(REGULAR_COLOR),
		BottomStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), false, true, true, true).
			Padding(0, 1).
			Foreground(INACTIVE_COLOR),
		card:     card,
		selected: selected,
	}
}

var ScrollBackground = InactiveStyle.Render(1, "┃")
var ScrollForeground = TextStyle.Render(1, "┃")
