package main

import (
	//	"fmt"

	"github.com/tamura2004/nobugo/ui"
)

var game *Game

type Game struct {
	Turn int
	Step Step
}

func init() {
	game = &Game{}
}

func (g *Game) Next() bool {
	if g.Turn == 0 {
		g.Turn++
		return g.SetUp()
	}

	switch g.Step {
	case PREPARE:
		g.Step = MARCH
		return true
	case MARCH:
		player.Println()
	case EMPLOY:
	case BATTLE:
	case CHECK:
	}

	return false
}

func (g *Game) SetUp() bool {
	ui.MsgBox("信長さんと部下の野望")
	ui.Pause("game start:")
	n := ui.InputNumber("プレイヤー人数を入力してください?", 3, 6)
	player.Num = n
	player.Active = 3
	for i := 1; i < n+1; i++ {
		samurai.Deck[i-1].Owner = Owner(i)
	}
	return true
}
