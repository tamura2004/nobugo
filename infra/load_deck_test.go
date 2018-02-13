package infra_test

import (
	"github.com/kr/pretty"
	"github.com/tamura2004/nobugo/infra"
)

func ExampleSamuraiDeckLoad() {
	deck := infra.LoadDeck("../config/samurai.toml")
	card := deck.Draw()
	pretty.Print(card)
	// Output:
	// domain.Card{
	//     Type:    0,
	//     Area:    0,
	//     Country: "",
	//     Name:    "弥助",
	//     Ability: 1,
	//     Power:   {1, 2, 6},
	// }
}

func ExampleCastleDeckLoad() {
	deck := infra.LoadDeck("../config/castle.toml")
	card := deck.Draw()
	pretty.Print(card)
	// Output:
	// domain.Card{
	//     Type:    0,
	//     Area:    1,
	//     Country: "羅馬",
	//     Name:    "法王庁",
	//     Ability: 0,
	//     Power:   {1, 1, 1},
	// }
}
