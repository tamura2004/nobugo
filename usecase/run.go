package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

var (
	step     Step
	gameOver bool
	Party    *domain.Party
	Board    *domain.Board
	Castle   domain.Deck
	Samurai  domain.Deck
)

func Run() {
	open()

	for !gameOver {
		step.next()
		Board.Print()
		Party.Print()
		UI.Pause("next")
	}
	close()
}
