package main

import (
	"io"
	"os"
	"os/exec"
)

// entry point where we can build tmux
// sessions.
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
