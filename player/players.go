package player

import (
	"fmt"
	"github.com/tamura2004/nobugo/game/board"
	"github.com/tamura2004/nobugo/table"
)

type Players []*Player

var pl Players

func Setup(num int) {
	pl = make(Players, num)
	for i := 0; i < num; i++ {
		pl[i] = New(i)
	}
}

func List() Players {
	return pl
}

func Prepare() {
	for i := 0; i < len(pl); i++ {
		pl[i].Prepare()
	}
}

func March() {
	fmt.Print("here?")

	for DiceRemain() {
		for i := 0; i < len(pl); i++ {
			pl[i].March()
		}
	}
}

func Employ() {
	for i := 0; i < 2; i++ {
		if w := board.Winer(i); w != -1 {
			pl[w].AddSamurai(board.GetSamurai(i))
		}
	}
}

func Battle() {
	for i := 2; i < 6; i++ {
		if w := board.Winer(i); w != -1 {
			pl[w].AddCastle(board.GetCastle(i))
		}
	}
}

func DiceRemain() bool {
	for i := 0; i < len(pl); i++ {
		if pl[i].Pool.Num > 0 {
			return true
		}
	}
	return false
}

func (ps Players) Table() (table []map[string][]string) {
	for i := 0; i < len(ps); i++ {
		table = append(table, ps[i].Value())
	}
	return table
}

func Print() {
	fmt.Println("\n---- プレイヤー一覧 ----")
	table.Render(pl, []string{"プレイヤー", "ダイス", "武将", "城"})
}
