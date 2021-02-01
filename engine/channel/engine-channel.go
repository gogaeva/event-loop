package engine_channel

import (
	. "../../iface"
)

type EventLoop struct {
	Handler
	queue  chan Command
	finish bool
}

func (loop *EventLoop) Start() {
	loop.queue = make(chan Command)
	go func() {
		for cmd := range loop.queue {
			cmd.Execute(loop)
			if len(loop.queue) == 0 && loop.finish {
				break
			}
		}
	}()
}

func (loop *EventLoop) Post(cmd Command) {
	loop.queue <- cmd
}

func (loop *EventLoop) AwaitFinish() {
	loop.finish = true
}
