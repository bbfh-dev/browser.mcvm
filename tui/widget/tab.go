package widget

import (
	"fmt"

	"github.com/bbfh-dev/browser.mcvm/tui/style"
)

const ITEMS_PER_PAGE = 20

type Tab interface {
	Normal() string
	Focused() string
}

type SimpleTab struct {
	id int
}

func NewSimpleTab(id int) SimpleTab {
	return SimpleTab{
		id: id,
	}
}

func (tab SimpleTab) Normal() string {
	return " "
}

func (tab SimpleTab) Focused() string {
	return " "
}

type TextTab struct {
	text string
}

func NewTextTab(text string) TextTab {
	return TextTab{
		text: text,
	}
}

func (tab TextTab) Normal() string {
	return fmt.Sprintf("[%s]", tab.text)
}

func (tab TextTab) Focused() string {
	return style.AccentStyle.Render(-1, tab.Normal())
}
