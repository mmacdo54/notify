package main

import (
	"fmt"
	"os/exec"
)

func wrapInQuotes(text string) string {
	return "\"" + text + "\""
}

type notification struct {
	title   string
	message string
	time    int64
}

func (n notification) GenerateCommand() string {
	return fmt.Sprintf("display notification %s with title %s", wrapInQuotes(n.message), wrapInQuotes(n.title))
}

func main() {
	notify(notification{
		title:   "Hey",
		message: "World",
		time:    1,
	})
}

func notify(n notification) {
	if err := exec.Command("osascript", "-e", n.GenerateCommand()).Run(); err != nil {
		fmt.Println(err)
	}
}
