package main

import (
	"fmt"
	"strconv"
)

type PrintCommand struct {
	arg string
}

func (pc PrintCommand) Execute(loop Handler) {
	fmt.Println(pc.arg)
}

type MulCommand struct {
	arg1, arg2 int
}

func (mc MulCommand) Execute(loop Handler) {
	res := mc.arg1 * mc.arg2
	loop.Post(&PrintCommand{arg: strconv.Itoa(res)})
}
