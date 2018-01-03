package game

//import "strings"
import (
	"fmt"
	"os"
	//	"math/rand"
	"github.com/olekukonko/tablewriter"
	"github.com/tamura2004/nobugo/deck"
	"github.com/tamura2004/nobugo/samurai"
)

type Phase int

const (
	Prepare Phase = iota
	March
	Employ
	Battle
	Finish
)

func (p Phase) String() string {
	switch p {
	case Prepare:
		return "準備フェイズ"
	case March:
		return "行軍フェイズ"
	case Employ:
		return "調略フェイズ"
	case Battle:
		return "合戦フェイズ"
	case Finish:
		return "終了フェイズ"
	default:
		return "不明なフェイズ"
	}
}

func (p *Phase) Next() {
	switch *p {
	case Prepare:
		*p = March
	case March:
		*p = Employ
	case Employ:
		*p = Battle
	case Battle:
		*p = Finish
	case Finish:
		*p = Prepare
	}
}

type Game struct {
	Phase
	turn          int
	currentPlayer int
	Players       deck.Deck
	jikisanDeck   deck.Deck
}

func (g *Game) Println() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ターン", "フェイズ", "手番プレイヤー"})
	table.Append([]string{
		fmt.Sprintf("%d", g.turn),
		g.Phase.String(),
		g.Players[0].Name(),
	})
	table.Render()
}

func (g *Game) NewPlayers(n int) {
	g.Players = g.jikisanDeck[:n]
}

func New() Game {
	g := Game{}
	g.turn = 1
	g.jikisanDeck = samurai.NewJikisanDeck()
	return g
}
