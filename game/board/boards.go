package board

import (
	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/game/board/box"
	. "github.com/tamura2004/nobugo/player/color"
	"github.com/tamura2004/nobugo/samurai"
	"github.com/tamura2004/nobugo/table"
	"strconv"
)

type Board struct {
	Box      [6]box.Box
	Name     [6]string
	No       [6]string
	Samurais [6]samurai.Samurai
	Castles  [6]castle.Castle
}

var board *Board

func Init() {
	no := [6]string{}
	name := [6]string{}
	samurais := [6]samurai.Samurai{}
	castles := [6]castle.Castle{}
	bx := [6]box.Box{}

	for i := 0; i < 6; i++ {
		no[i] = strconv.Itoa(i + 1)
		if i < 2 {
			name[i] = "行軍"
		} else {
			name[i] = "合戦"
		}
		bx[i] = box.New()
	}

	board = &Board{
		No:       no,
		Name:     name,
		Samurais: samurais,
		Castles:  castles,
		Box:      bx,
	}
}

func Stat() *Board {
	if board != nil {
		return board
	}
	Init()
	return board
}

func Prepare() {
	for i := 0; i < 6; i++ {
		board.Box[i] = box.New()

		if i < 2 {
			s, err := samurai.DrawOne()
			if err == nil {
				board.Samurais[i] = s
			}
		} else {
			c, err := castle.DrawOne()
			if err == nil {
				board.Castles[i] = c
			}
		}
	}
}

func Winer(i int) Color {
	return board.Box[i].Winer()
}

func GetSamurai(i int) samurai.Samurai {
	return board.Samurais[i]
}

func GetCastle(i int) castle.Castle {
	return board.Castles[i]
}

func Print() {
	table.Render(board, []string{"アクション", "カード", "ダイス"})
}

func (b Board) Row(i int) map[string][]string {
	var v []string
	if i < 2 {
		v = b.Samurais[i].Value()
	} else {
		v = b.Castles[i].Value()
	}
	return map[string][]string{
		"アクション": {b.No[i], b.Name[i]},
		"カード":   v,
		"ダイス":   b.Box[i].Value(),
	}
}

func (b Board) Table() []map[string][]string {
	v := [6]map[string][]string{}
	for i := 0; i < 6; i++ {
		v[i] = b.Row(i)
	}
	return v[:]
}
