package tui

import "github.com/bbfh-dev/browser.mcvm/tui/widget"

var KEYBINDS = map[string]widget.Keybind{
	"quit":         widget.NewKeybind("q", "ctrl+c"),
	"goto.top":     widget.NewKeybind("g"),
	"goto.bottom":  widget.NewKeybind("G"),
	"scroll.up":    widget.NewKeybind("down", "J"),
	"scroll.down":  widget.NewKeybind("up", "K"),
	"tab.previous": widget.NewKeybind("left", "H"),
	"tab.next":     widget.NewKeybind("right", "L"),
	"search":       widget.NewKeybind("/"),
}
