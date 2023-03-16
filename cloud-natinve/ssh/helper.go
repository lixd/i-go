package sshutils

import (
	"fmt"
	"strings"
)

func withSudo(sshConfig *SSH, cmd string) string {
	if sshConfig.User == "root" {
		return cmd
	}
	// NOTE: spilt cmd by &&,not work on some command,such as 'echo "&&"'
	split := strings.Split(cmd, "&&")
	list := make([]string, 0, len(split))
	for _, v := range split {
		var sudoCmd string
		if sshConfig.Password != "" {
			// use `echo '$passwd' |sudo -S $cmd` cmd avoid interactive enter passwd
			sudoCmd = fmt.Sprintf("echo '%s' | sudo -S %s", sshConfig.Password, v)
		} else {
			// no passwd maybe user configured NOPASSWD in sudoers.
			sudoCmd = "sudo " + v
		}
		list = append(list, sudoCmd)
	}
	return strings.Join(list, "&&")
}

// printCmd replace sensitive information in cmd
func printCmd(passwd, cmd string) string {
	if passwd == "" {
		return cmd
	}
	return strings.ReplaceAll(cmd, fmt.Sprintf("echo '%s'", passwd), "echo $PASSWD")
}
