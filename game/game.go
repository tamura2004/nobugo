package game

import (
	//	"fmt"

	"github.com/tamura2004/nobugo/player"

	"github.com/tamura2004/nobugo/board"
	"github.com/tamura2004/nobugo/step"
	"github.com/tamura2004/nobugo/ui"
)

var game *Game

type Game struct {
	Player *Player
	Turn   int
	Step   step.Step
}

func init() {
	game = new(Game)
	game.Turn = 1
}

func (g *Game) Next() bool {
	switch g.Step {
	case PREPARE:
		for _, p := range player.Deck {
			p.Pool.Num = 3
		}
		SetupActionSamurai()
		SetupActionCastle()
		g.Step = MARCH
		return true
	case MARCH:
		PrintPlayer()
		PrintAction()
	case EMPLOY:
	case BATTLE:
	case CHECK:
	}

	return false
}
