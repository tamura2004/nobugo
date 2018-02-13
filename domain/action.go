package domain

type Action struct {
	Msg string
	Do  func()
}
