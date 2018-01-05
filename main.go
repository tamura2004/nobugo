package main

//import "strings"
import (
	"fmt"
	//	"os"
	//	"math/rand"
	//	"github.com/abiosoft/ishell"
	//	"github.com/olekukonko/tablewriter"
	//	"github.com/tamura2004/nobugo/deck"
	"github.com/tamura2004/nobugo/samurai"
)

func main() {
	deck := samurai.Deck()
	deck.Shuffle()
	fmt.Printf("%V", deck)
	/*
			g := game.New()
		//	shell := ishell.New()

			fmt.Println("# 信長さんの野望")

			var n int
			for n < 3 || n > 6 {
				fmt.Print("プレイヤー人数?(3-6)> ")
				fmt.Scan(&n)
			}
			g.NewPlayers(n)

			g.Println()

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"手番", "名前", "ダイス"})

			data := [][]string{
				[]string{"レ", g.Players[0].Name(), "[1,2,3]"},
				[]string{"", g.Players[1].Name(), "[2,3,4]"},
				[]string{"", g.Players[2].Name(), "[5,6,6]"},
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
					for _, player := range g.Players {
						c.Println(player.Name())
					}
				},
			})

			shell.Run()
	*/
}
