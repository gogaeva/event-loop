package parser

import (
	"strconv"
	"strings"

	. "../commands"
	. "../iface"
)

func Parse(cmdLine string) (cmd Command) {
	parts := strings.Fields(cmdLine)
	cmdName := parts[0]
	switch cmdName {
	case "print":
		cmd = PrintCommand{parts[1]}
	case "mul":
		arg1, _ := strconv.Atoi(parts[1])
		arg2, _ := strconv.Atoi(parts[2])
		cmd = MulCommand{arg1, arg2}
	}
	return
}
