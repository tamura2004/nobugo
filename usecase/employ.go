package usecase

import (
	"github.com/kr/pretty"
)

func employ() {
	for i := 0; i < 2; i++ {
		pool := Board[i].Pool
		pretty.Print(pool)

		color := pool.Winner()
		pretty.Print(color)

		pl := Party.GetPlayerByColor(color)
		pretty.Print(pl)

		card := Board[i].Deck.Draw()
		pretty.Print(card)

		pl.Samurai.Bottom(card)
		pretty.Print(pl)

	}
}
