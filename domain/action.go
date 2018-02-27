package domain

type Action struct {
	Msg string
	Do  func()
}

var Nop Action = Action{
	Msg: "no action",
	Do:  func() {},
}
