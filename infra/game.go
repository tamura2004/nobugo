package infra

import (
	"fmt"
	"github.com/tamura2004/nobugo/adapter"
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/usecase"
)

func NewGame() *usecase.Game {
	n := 3
	g := &usecase.Game{
		Game:  domain.NewGame(),
		Party: usecase.NewParty(n),
		Board: usecase.NewBoard(),
	}

	g.Party.Printer = adapter.NewPartyPrinter(&g.Party.Party)
	g.Board.Printer = adapter.NewPartyPrinter(g.Board)

	return g
}
