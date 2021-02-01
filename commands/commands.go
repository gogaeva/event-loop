package commands

import (
	"fmt"
	"strconv"

	. "../iface"
)

type PrintCommand struct {
	Arg string
}

func (pc PrintCommand) Execute(loop Handler) {
	fmt.Println(pc.Arg)
}

type MulCommand struct {
	Arg1, Arg2 int
}

func (mc MulCommand) Execute(loop Handler) {
	res := mc.Arg1 * mc.Arg2
	go loop.Post(&PrintCommand{Arg: strconv.Itoa(res)})
}
