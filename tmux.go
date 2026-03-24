package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

// entry point where we can build tmux
// commands.
func runTmux(args ...string) error {
	cmd := exec.Command("tmux", args...)

	// looks like we need to set
	// at least cmd.Stdin to os.Stdin,
	// otherwise this goes to /dev/null
	// and we get an EOF
	cmd.Stdin = os.Stdin
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run()
}

// actual attach function. we should only
// attach the session when all the commands have been
// completed succesfully to prevent flickering and
// bleeding to stdout
func Attach(session string) error {
	tmuxPath, err := exec.LookPath("tmux")
	if err != nil {
		return err
	}
	return syscall.Exec(
		tmuxPath,
		[]string{"tmux", "attach", "-t", session},
		os.Environ(),
	)
}

// TmuxSession queues up tmux commands and executes them
// all at once when Launch is called.
type TmuxSession struct {
	name       string
	commands   [][]string
	hasWindows bool
}

// NewTmuxSession creates a new session builder. The session
// creation is deferred until the first AddWindow call so that
// the first window's command can be baked into new-session,
// avoiding an extra empty shell window.
func NewTmuxSession(name string) *TmuxSession {
	return &TmuxSession{
		name: name,
	}
}

// AddWindow queues a new tmux window. The first call creates
// the session itself with this window's command, so there is
// no leftover empty shell window. Subsequent calls use
// new-window.
func (s *TmuxSession) AddWindow(name string, command string) {
	if !s.hasWindows {
		// First window: create the session with this command
		args := []string{"new-session", "-s", s.name, "-n", name, "-d"}
		if command != "" {
			args = append(args, command)
		}
		s.commands = append(s.commands, args)
		s.hasWindows = true
	} else {
		args := []string{"new-window", "-t", s.name, "-n", name}
		if command != "" {
			args = append(args, command)
		}
		s.commands = append(s.commands, args)
	}
}

func (s *TmuxSession) Launch() error {
	for _, args := range s.commands {
		if err := runTmux(args...); err != nil {
			_ = runTmux("kill-session", "-t", s.name)
			return fmt.Errorf("tmux %v: %w", args, err)
		}
	}
	return Attach(s.name)
}
