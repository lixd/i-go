package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("git")
	if lookErr != nil {
		panic(lookErr)
	}

	// args := []string{"ls", "-a", "-l", "-h"}
	args := []string{"git", "-v"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
