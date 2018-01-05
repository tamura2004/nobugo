package game_test

import (
	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/deck"
	"github.com/tamura2004/nobugo/game"
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
	// --
}
