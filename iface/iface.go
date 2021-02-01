package iface

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}
