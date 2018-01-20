package board

import (
	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/game/board/box"
	"github.com/tamura2004/nobugo/samurai"
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

func New() Board {
	sd := samurai.Deck()
	cd := castle.Deck()

	no := [6]string{}
	name := [6]string{}
	card := [6]Card{}
	bx := [6]box.Box{}

	for i := 0; i < 6; i++ {
		no[i] = strconv.Itoa(i + 1)
		if i < 2 {
			name[i] = "行軍"
			card[i] = sd[i]
		} else {
			name[i] = "合戦"
			card[i] = cd[i-2]
		}
		bx[i] = box.New()
	}

	return Board{
		No:   no,
		Name: name,
		Card: card,
		Box:  bx,
	}
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
