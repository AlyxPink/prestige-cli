package main

import (
	"log"
	"os"

	"github.com/AlyxPink/prestige-cli/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			log.Fatal("fatal:", err)
		}
		defer f.Close()
	}
	p := tea.NewProgram(
		ui.NewModel(),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
