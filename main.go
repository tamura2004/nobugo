package main

//import "strings"
import (
	"fmt"
	"os"
	//	"math/rand"
	"github.com/abiosoft/ishell"
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
	players       deck.Deck
	jikisanDeck   deck.Deck
}

func main() {
	g := NewGame()
	shell := ishell.New()

	shell.Println("信長さんの野望")
	shell.Println("\n")

	var n int
	for n < 3 || n > 6 {
		fmt.Print("プレイヤー人数?(3-6): ")
		fmt.Scan(&n)
	}
	g.players = g.jikisanDeck[:n]

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"手番", "名前", "ダイス"})

	data := [][]string{
		[]string{"レ", g.players[0].Name(), "[1,2,3]"},
		[]string{"", g.players[1].Name(), "[2,3,4]"},
		[]string{"", g.players[2].Name(), "[5,6,6]"},
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	shell.AddCmd(&ishell.Cmd{
		Name: "n",
		Help: "done action, and progress next phase.",
		Func: func(c *ishell.Context) {
			g.Next()
			c.Println(g.String())
			for _, player := range g.players {
				c.Println(player.Name())
			}
		},
	})

	shell.Run()
}

func NewGame() Game {
	g := Game{}
	g.jikisanDeck = samurai.NewJikisanDeck()
	return g
}
