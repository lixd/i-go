package sshutils

import (
	"fmt"
	"strings"
)

type Result struct {
	User     string
	Host     string
	Cmd      string
	PrintCmd string
	Stdout   string
	Stderr   string
	ExitCode int
}

func (r Result) Short() string {
	return fmt.Sprintf("(run `%s` on %s@%s)", r.PrintCmd, r.User, r.Host)
}

func (r Result) Error() error {
	if r.ExitCode == 0 {
		return nil
	}
	return fmt.Errorf("%s err: %s", r.Short(), r.Stderr)
}

func (r Result) String() string {
	return fmt.Sprintf("(run `%s` on %s@%s) \nstdout: \n%s\nstderr: %s\nexitcode: %v\n", r.PrintCmd, r.User, r.Host,
		r.Stdout, r.Stderr, r.ExitCode)
}

func (r Result) StdoutToString(place string) string {
	return strings.ReplaceAll(r.Stdout, "\n", place)
}
