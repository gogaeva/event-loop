package main

import (
	"bufio"
	"flag"
	"os"

	engine "./engine/channel"
	parser "./parser"
)

func main() {
	inputFile := flag.String("f", "", "input file name")
	flag.Parse()
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(*inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parser.Parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
