package usecase

import (
	"github.com/kr/pretty"
)

func employ() {
	for i := 0; i < 2; i++ {
		pool := Board[i].Pool
		color := pool.Winner()
		pl := Party.GetPlayerByColor(color)
		card := Board[i].Deck.Draw()
		pl.Samurai.Bottom(card)

		pretty.Print(pool)
		pretty.Print(color)
		pretty.Print(pl)
		pretty.Print(card)
	}
}
