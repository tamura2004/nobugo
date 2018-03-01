package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

func employ() {
	for i := 0; i < 2; i++ {
		box := &Board[i]
		card := box.Deck.Draw()
		color := box.Pool.Winner()

		box.Pool = domain.NewEmptyPool()

		pl := Party.GetPlayerByColor(color)
		pl.Samurai.Bottom(card)
	}
}
