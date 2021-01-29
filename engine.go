package main

type EventLoop struct {
	Handler
	queue chan Command
	//head, tail int
	empty chan bool
}

// type messageQueue struct {

// }

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

func (loop *EventLoop) Start() {
	loop.queue = make(chan Command, 10)
	loop.empty = make(chan bool)
	go func() {
		for cmd := range loop.queue {
			cmd.Execute(loop)
			if len(loop.queue) == 0 {
				loop.empty <- true
			}
		}
	}()
}

func (loop *EventLoop) Post(cmd Command) {
	loop.queue <- cmd
}

func (loop *EventLoop) AwaitFinish() {
	for {
		if <-loop.empty {
			close(loop.queue)
			return
		}
	}
}
