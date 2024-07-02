package style

import (
	"github.com/charmbracelet/lipgloss"
)

type Style interface {
	Render(width int, content string) string
	RenderLine(width int, content string) string
}

type defaultStyle struct {
	Style lipgloss.Style
}

func (text defaultStyle) Render(width int, content string) string {
	if width == -1 {
		return text.Style.Render(content)
	}
	return text.Style.Width(width).Render(content)
}

func (text defaultStyle) RenderLine(width int, content string) string {
	return text.Render(width, content) + "\n"
}

type Card interface {
	Name() string
	Description() string
	Hint() string
	Valid() bool
	Extra() string
}

type cardStyle struct {
	TopStyle    lipgloss.Style
	MiddleStyle lipgloss.Style
	BottomStyle lipgloss.Style
	card        Card
	selected    bool
}

func (card cardStyle) Render(width int, _ string) string {
	if card.selected {
		card.TopStyle = card.TopStyle.BorderForeground(ACCENT_COLOR).Foreground(ACCENT_COLOR)
		card.MiddleStyle = card.MiddleStyle.BorderForeground(ACCENT_COLOR)
		card.BottomStyle = card.BottomStyle.BorderForeground(ACCENT_COLOR)
	}

	var middle string
	if card.card.Valid() {
		middle = card.MiddleStyle.Width(width - 2).Render(card.card.Description())
	} else {
		middle = card.MiddleStyle.Width(width - 2).Foreground(ERROR_COLOR).Render("Couldn't load the package data!")
	}

	extra := card.card.Extra()
	if lipgloss.Width(extra) >= width-3 {
		extra = extra[:width-3] + "…"
	}

	return lipgloss.JoinVertical(
		0,
		card.TopStyle.Width(width-2).
			Render(lipgloss.JoinHorizontal(0, WithIcon(PACK_ICON, card.card.Name()), InactiveStyle.Render(-1, " #"+card.card.Hint()))),
		middle,
		card.MiddleStyle.Width(width-2).Render(""),
		card.BottomStyle.Width(width-2).Render(extra),
	)
}

func (text cardStyle) RenderLine(width int, content string) string {
	return text.Render(width, content) + "\n"
}
