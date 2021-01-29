package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := flag.String("f", "", "input file name")
	flag.Parse()
	eventLoop := new(EventLoop)
	eventLoop.Start()
	if input, err := os.Open(*inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}

func parse(cmdLine string) (cmd Command) {
	parts := strings.Fields(cmdLine)
	cmdName := parts[0]
	switch cmdName {
	case "print":
		cmd = PrintCommand{arg: parts[1]}
	case "mul":
		arg1, _ := strconv.Atoi(parts[1])
		arg2, _ := strconv.Atoi(parts[2])
		cmd = MulCommand{arg1, arg2}
	}
	return
}
