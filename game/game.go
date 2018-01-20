package game

import (
	//	"fmt"

	"github.com/tamura2004/nobugo/player"

	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/game/board"
	. "github.com/tamura2004/nobugo/game/step"
	"github.com/tamura2004/nobugo/samurai"
	"github.com/tamura2004/nobugo/ui"
)

var game Game

type Game struct {
	Turn int
	Step Step
	player.Players
	samurai.Samurais
	castle.Castles
	board.Board
}

func init() {
	game = Game{
		Turn:     1,
		Step:     PREPARE,
		Samurais: samurai.Deck(),
		Castles:  castle.Deck(),
		Board:    board.New(),
	}
}

func New() (g *Game) {
	num := ui.InputNumber("プレイヤー人数？", 3, 6)
	game.Players = player.Setup(num)
	return &game
}

func (g *Game) Next() bool {
	switch g.Step {
	case PREPARE:
		for i := 0; i < len(g.Players); i++ {
			g.Players[i].Prepare()
		}
		g.Board.Prepare()
		g.Step = MARCH
		return true
	case MARCH:
		table.Render(g.Playres, []string{"プレイヤー", "ダイス", "武将", "城"})
		table.Render(g.Board, []string{"アクション", "カード", "ダイス"})
	case EMPLOY:
	case BATTLE:
	case CHECK:
	}

	return false
}
