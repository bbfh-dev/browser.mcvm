package widget

import (
	"strings"

	"github.com/bbfh-dev/browser.mcvm/pkg/enum"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchWidget struct {
	Text    string
	Focused bool
}

func NewSearchWidget() SearchWidget {
	return SearchWidget{
		Text:    "",
		Focused: false,
	}
}

func (widget SearchWidget) Blur() SearchWidget {
	widget.Focused = false
	return widget
}

func (widget SearchWidget) Focus() SearchWidget {
	widget.Focused = true
	return widget
}

func HasInput(input string) bool {
	return len(input) != 0
}

func (widget SearchWidget) Empty() bool {
	return !HasInput(widget.Text)
}

func (widget SearchWidget) HandleTyping(msg tea.KeyMsg) SearchWidget {
	switch msg.String() {
	case "esc":
		widget.Text = ""
		widget = widget.Blur()
	case "enter":
		widget = widget.Blur()
	case "backspace":
		if len(widget.Text) > 0 {
			widget.Text = widget.Text[:len(widget.Text)-1]
		}
	case "ctrl+h":
		widget.Text = ""
	default:
		if strings.Contains(enum.INPUT_CHARACTERS, msg.String()) {
			widget.Text += msg.String()
		}
	}
	return widget
}
