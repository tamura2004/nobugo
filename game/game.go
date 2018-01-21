package game

import (
	"fmt"
	"github.com/tamura2004/nobugo/player"

	"github.com/tamura2004/nobugo/game/board"
	. "github.com/tamura2004/nobugo/game/step"
	"github.com/tamura2004/nobugo/ui"
)

var game Game

type Game struct {
	Turn int
	Step Step
	player.Players
	board.Board
}

func init() {
	game = Game{
		Turn: 1,
		Step: PREPARE,
	}
}

func New() (g *Game) {
	board.Init()
	num := ui.InputNumber("プレイヤー人数？", 3, 6)
	player.Setup(num)
	return &game
}

func (g *Game) Next() bool {
	switch g.Step {
	case PREPARE:
		player.Prepare()
		board.Prepare()
		g.Step = MARCH
		return true
	case MARCH:
		player.Print()
		board.Print()
		player.March()
		g.Step = EMPLOY
		return true
	case EMPLOY:
		player.Employ()
		g.Step = BATTLE
		return true
	case BATTLE:
		player.Battle()
		g.Step = CHECK
		return true
	case CHECK:
		g.Turn++
		fmt.Printf("\n\n---- %d turn end ----\n\n", g.Turn)
		ui.Pause("press enter >")
		g.Step = PREPARE
		return true
	}

	return false
}
