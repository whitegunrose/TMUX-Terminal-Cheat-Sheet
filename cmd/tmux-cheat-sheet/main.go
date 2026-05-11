package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/whitegunrose/TMUX-Terminal-Cheat-Sheet/internal/ui"
)

func main() {
	// fmt.Println("어서 오세요!")

	p := tea.NewProgram(
		ui.New(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
