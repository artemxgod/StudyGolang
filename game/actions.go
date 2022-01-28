
package main

import (
	"fmt"
	"strings"
)

func handleCommand(cmd string) string{
	actions := strings.Split(cmd, " ")
	for _, v := range actions{
		fmt.Print(translate(v))
	}
	return ""
}

func translate(cmd string) string {
	list := map[string]string{
		"осмотреться": "lookup",
		"идти":        "go",
		"взять":       "take",
		"применить":   "use",
		"одеть":       "clothe",
	}
	return list[cmd]
}