package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

type Game struct {
	domain.Game
}

func (g *Game) Run() {
	for g.next() {
		g.Party.Print()
		g.Board.Print()
	}
	g.close()
}

func (g *Game) close() {
}

func (g *Game) next() bool {
	switch g.Step {

	case domain.PREPARE:
		g.prepare()

	case domain.MARCH:
		g.march()

	case domain.EMPLOY:
		g.employ()

	case domain.BATTLE:
		g.battle()

	case domain.CHECK:
		g.check()
	}
	return false
}

func (g *Game) prepare() {
}

func (g *Game) march() {
}

func (g *Game) employ() {
}

func (g *Game) battle() {
}

func (g *Game) check() {
}
