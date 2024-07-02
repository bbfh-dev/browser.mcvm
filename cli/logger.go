package cli

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// Write information into the log file.
// Uses fmt.Sprintf() format
func Log(format string, v ...any) {
	log.Printf(format, v...)
}

// Sets up logger output file and flags.
// Returns pointer to the log file that should be used
// to `defer file.Close()`
func SetupLogger() *os.File {
	file, err := tea.LogToFile(filepath.Join(os.TempDir(), "browse_mcvm.log"), "")
	if err != nil {
		fmt.Println("FATAL: ", err)
		os.Exit(1)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime)

	return file
}
