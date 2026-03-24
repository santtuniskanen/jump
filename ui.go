package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.ANSIColor(5)).MarginBottom(1)
	itemStyle     = lipgloss.NewStyle().PaddingLeft(2)
	selectedStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.ANSIColor(6)).Bold(true)
)

type entry struct {
	name    string
	host    string
	command string
}

type model struct {
	entries  []entry
	cursor   int
	chosen   *entry
	quitting bool
}

func newModel(hosts []Host) model {
	entries := make([]entry, 0, len(hosts))
	for _, h := range hosts {
		cmd := fmt.Sprintf("ssh -q -t %s@%s", h.User, h.Address)
		if h.Key != "" {
			cmd = fmt.Sprintf("ssh -q -t -i %s %s@%s", h.Key, h.User, h.Address)
		}
		entries = append(entries, entry{
			name:    h.Name,
			host:    h.Address,
			command: cmd,
		})
	}

	if len(entries) == 0 {
		log.Println("no entries found")
		os.Exit(1)
	}

	return model{
		entries: entries,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
			}
		case "enter":
			chosen := m.entries[m.cursor]
			m.chosen = &chosen
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var s strings.Builder

	s.WriteString(titleStyle.Render("select a host") + "\n")

	for i, e := range m.entries {
		line := fmt.Sprintf("%s (%s)", e.name, e.host)
		if i == m.cursor {
			s.WriteString(selectedStyle.Render("> "+line) + "\n")
		} else {
			s.WriteString(itemStyle.Render("  "+line) + "\n")
		}
	}
	return s.String()
}
