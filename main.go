package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	config, err := loadConfig("config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config: %s\n", err.Error())
	}

	p := tea.NewProgram(newModel(config.Hosts))
	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running TUI: %v\n", err)
		os.Exit(1)
	}

	m := finalModel.(model)
	if m.chosen == nil {
		os.Exit(0)
	}

	session := NewTmuxSession(m.chosen.name)

	if m.chosen.command != "" {
		session.AddWindow("shell", m.chosen.command)
	}

	if err := session.Launch(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
