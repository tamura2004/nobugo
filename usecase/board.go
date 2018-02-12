package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

type Board struct {
	domain.Board
	Printer
}

func NewBoard() *Board {
	return &Board{
		Board: domain.NewBoard(),
	}
}

type Box struct {
	domain.Box
}
