package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Player struct {
	Color Color
	Owner Owner
	Name  string
	Gold  int
}

var player struct {
	Num     int
	Active  int
	Deck    []*Player
	Println func()
}

func init() {
	player.Deck = make([]*Player, 7)
	for i := 1; i <= 6; i++ {
		player.Deck[i] = NewPlayer(i)
	}

	player.Println = PrintPlayer
}

func PrintPlayer() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetRowLine(true)
	t.SetAutoMergeCells(true)

	data := make([][]string, 3)

	for i := 0; i < len(data); i++ {
		data[i] = make([]string, player.Num+1)
	}

	data[0][0] = "プレイヤー"
	data[1][0] = "ダイス"
	data[2][0] = "武将"

	aligns := make([]int, player.Num+1)
	aligns[0] = tablewriter.ALIGN_DEFAULT

	for i := 1; i <= player.Num; i++ {
		t.SetColMinWidth(i, 10)
		aligns[i] = tablewriter.ALIGN_CENTER
		c := Color(i).String()
		if i == player.Active {
			data[0][i] = "手番 -> " + c
		} else {
			data[0][i] = c
		}
		data[1][i] = "3"
		data[2][i] = samurai.Deck[i-1].Name
	}

	t.SetColumnAlignment(aligns)
	t.AppendBulk(data)
	t.Render()
}

func NewPlayer(i int) *Player {
	return &Player{
		Color: Color(i),
		Owner: Owner(i),
	}
}
