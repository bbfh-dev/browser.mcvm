package style

import "github.com/charmbracelet/lipgloss"

type Style interface {
	Render(width int, content string) string
	RenderLine(width int, content string) string
}

type defaultStyle struct {
	style lipgloss.Style
}

func (text defaultStyle) Render(width int, content string) string {
	if width == -1 {
		return text.style.Render(content)
	}
	return text.style.Width(width).Render(content)
}

func (text defaultStyle) RenderLine(width int, content string) string {
	return text.Render(width, content) + "\n"
}

type quoteStyle struct {
	defaultStyle
	icon string
}
