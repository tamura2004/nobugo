package game

//import "strings"
import (
	"fmt"
	"os"
	//	"math/rand"
	"github.com/olekukonko/tablewriter"
	"github.com/tamura2004/nobugo/deck"
)

type Phase int

const (
	Prepare Phase = iota + 1
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
	turn int
	Phase
	Players     deck.Deck
	SamuraiDeck deck.Deck
	CastleDeck  deck.Deck
}

func (g *Game) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ターン", "フェイズ", "手番プレイヤー"})
	table.Append([]string{
		fmt.Sprintf("%d", g.turn),
		g.Phase.String(),
		g.SamuraiDeck[0].String(),
	})
	table.Render()
}

func New() Game {
	return Game{
		Phase: Prepare,
		turn:  1,
	}
}
