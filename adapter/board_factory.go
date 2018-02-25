package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type boardFactory func() *domain.Board

func (f boardFactory) Create() *domain.Board {
	return f()
}

func NewBoard() *domain.Board {
	b := &domain.Board{}
	for i := 0; i < 6; i++ {
		b.Box[i] = domain.NewBox(i)
	}
	b.Printer = BoardPresenter{b}
	return b
}
