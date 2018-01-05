package player_test

import (
	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/deck"
	"github.com/tamura2004/nobugo/player"
	"github.com/tamura2004/nobugo/samurai"
)

func ExamplePrintTable() {
	s := samurai.Deck()
	c := castle.Deck()
	x := deck.Deck{
		player.Player{true, s[0].String(), s[:3], c[:3], []int{1, 2, 3}},
		player.Player{false, s[1].String(), s[3:6], c[3:6], []int{1, 2, 3}},
	}
	x.PrintTable()
	// Output:
	//+------+------------+------------------------------+--------------------------------+---------+
	//| 手番 |    名前    |             武将             |               城               | ダイス  |
	//+------+------------+------------------------------+--------------------------------+---------+
	//| レ   | 木下藤吉郎 | 木下藤吉郎;明智光秀;柴田勝家 | 法王庁;東インド会社;万里の長城 | [1 2 3] |
	//|      | 明智光秀   | 前田利家;滝川一益;弥助       | 首里城;鹿児島城;高知城         | [1 2 3] |
	//+------+------------+------------------------------+--------------------------------+---------+
}
