package domain

type Board struct {
	Box [6]Box
	Printer
}

var BoardFactory interface {
	Create() *Board
}
