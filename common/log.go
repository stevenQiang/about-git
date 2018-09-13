package common

import (
	"os/exec"
	"strings"
)

// 因为git2go不支持git log
func Log(name string) int{
	result,err := exec.Command("git", "log", "--pretty=format:'%h'", "--author="+name).Output()
	CheckIfError(err)
	commit := strings.Split(string(result), "\n")
	return len(commit)
}
