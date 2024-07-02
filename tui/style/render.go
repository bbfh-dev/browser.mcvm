package style

import "github.com/charmbracelet/lipgloss"

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
	Valid() bool
}

type cardStyle struct {
	TopStyle    lipgloss.Style
	BottomStyle lipgloss.Style
	card        Card
	selected    bool
}

func (card cardStyle) Render(width int, _ string) string {
	if card.selected {
		card.TopStyle = card.TopStyle.BorderForeground(ACCENT_COLOR).Foreground(ACCENT_COLOR)
		card.BottomStyle = card.BottomStyle.BorderForeground(ACCENT_COLOR)
	}

	var bottom string
	if card.card.Valid() {
		bottom = card.BottomStyle.Width(width - 2).Render(card.card.Description())
	} else {
		bottom = card.BottomStyle.Width(width - 2).Foreground(ERROR_COLOR).Render("Couldn't load the package data!")
	}

	return lipgloss.JoinVertical(
		0,
		card.TopStyle.Width(width-2).Render(WithIcon(PACK_ICON, card.card.Name())),
		bottom,
	)
}

func (text cardStyle) RenderLine(width int, content string) string {
	return text.Render(width, content) + "\n"
}
