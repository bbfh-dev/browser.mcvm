package cli

import (
	"os"

	"github.com/jessevdk/go-flags"
)

// CLI Flags passed to the program when called
var Flags struct {
	Textonly bool `short:"t" long:"text-only" description:"Disable icons in TUI"`
}

// Parses the CLI flags and sets default values
func ParseFlags() {
	_, err := flags.Parse(&Flags)
	if err != nil {
		os.Exit(0)
	}
}
