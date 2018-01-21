package board

import (
	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/game/board/box"
	"github.com/tamura2004/nobugo/samurai"
	"github.com/tamura2004/nobugo/table"
	"strconv"
)

type Board struct {
	Box  [6]box.Box
	Name [6]string
	No   [6]string
	Card [6]Card
}

type Card interface {
	Value() []string
}

var board *Board

func Init() {
	no := [6]string{}
	name := [6]string{}
	card := [6]Card{}
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
		No:   no,
		Name: name,
		Card: card,
		Box:  bx,
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
		if i < 2 {
			board.Card[i] = samurai.Deck().Draw(1)[0]
		} else {
			board.Card[i] = castle.Deck().Draw(1)[0]
		}
	}
}

func Print() {
	table.Render(board, []string{"アクション", "カード", "ダイス"})
}

func (b Board) Row(i int) map[string][]string {
	return map[string][]string{
		"アクション": {b.No[i], b.Name[i]},
		"カード":   b.Card[i].Value(),
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
