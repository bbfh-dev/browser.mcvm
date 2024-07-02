package style

import (
	"fmt"

	"github.com/bbfh-dev/browser.mcvm/cli"
)

type icon string

// https://www.nerdfonts.com/cheat-sheet
const (
	GAME_ICON     icon = "󰍳 "
	LIST_ICON     icon = "󱉯 "
	DATABASE_ICON icon = " "
	PACK_ICON     icon = "󰆦 "
	SEARCH_ICON   icon = " "
)

// Prepends an icon if --icons CLI flag is on
func WithIcon(icon icon, str string) string {
	if !cli.Flags.Textonly {
		return fmt.Sprintf("%s %s", icon, str)
	}

	return str
}

// Returns an icon if --icons CLI flag is on, otherwise the fallback string
func IconFallback(icon icon, fallback string) string {
	if !cli.Flags.Textonly {
		return string(icon)
	}

	return fallback
}
