package style

import "github.com/charmbracelet/lipgloss"

var (
	SEARCH_COLOR   = lipgloss.Color("#9D448A")
	ACCENT_COLOR   = lipgloss.Color("#F36E2A")
	REGULAR_COLOR  = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}
	HINT_COLOR     = lipgloss.AdaptiveColor{Light: "#4E4E4E", Dark: "#969696"}
	INACTIVE_COLOR = lipgloss.AdaptiveColor{Light: "#A7A7A7", Dark: "#5E5E5E"}
	ERROR_COLOR    = lipgloss.AdaptiveColor{Light: "#DC2626", Dark: "#DC2626"}
)
