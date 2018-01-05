package deck_test

import (
	"fmt"
	"testing"

	"github.com/tamura2004/nobugo/deck"
	"github.com/tamura2004/nobugo/samurai"
)

func TestShuffle(t *testing.T) {
	deck := samurai.Deck()

	before := deck[0]
	deck.Shuffle()
	after := deck[0]

	if before == after {
		t.Errorf("Shuffle: before %v after %v is not changed", before, after)
	}
}

func ExampleShuffle() {
	x := samurai.Deck()
	x.Shuffle()
	fmt.Println(len(x))
	// Output:
	// 19
}

func ExampleDrawIf() {
	x := samurai.Deck()
	a := x.DrawIf(func(c deck.Card) bool {
		for _, v := range c.Values() {
			if v == "〇" {
				return true
			}
		}
		return false
	})
	fmt.Printf("%v;%v", a.Values(), len(x))
	// Output:
	// [0 木下藤吉郎 〇 [3]/[4]->[1]];18
}

func ExamplePrintTable() {
	x := samurai.Deck()
	x.PrintTable()
	// Output:
	//+------+--------------+------+--------------+
	//| 番号 |     名前     | 直臣 |     能力     |
	//+------+--------------+------+--------------+
	//|    0 | 木下藤吉郎   | 〇   | [3]/[4]->[1] |
	//|    1 | 明智光秀     | 〇   | [5]/[6]->[2] |
	//|    2 | 柴田勝家     | 〇   | [3]/[4]->[5] |
	//|    3 | 前田利家     | 〇   | [5]/[6]->[3] |
	//|    4 | 滝川一益     | 〇   | [1]/[2]->[4] |
	//|    5 | 弥助         | 〇   | [1]/[2]->[6] |
	//|    6 | 毛利元就     |      | [5]/[6]->[4] |
	//|    7 | 徳川家康     |      | [3]/[4]->[2] |
	//|    8 | 今川義元     |      | [5]/[6]->[1] |
	//|    9 | 北条氏康     |      | [1]/[2]->[3] |
	//|   10 | 島津義弘     |      | [1]/[2]->[5] |
	//|   11 | 松永久秀     |      | 交換：武将   |
	//|   12 | 武田信玄     |      | 城主：Ｄ＋１ |
	//|   13 | 上杉謙信     |      | 合戦：＋２Ｄ |
	//|   14 | 伊達政宗     |      | 東４         |
	//|   15 | 九鬼義隆     |      | 東６         |
	//|   16 | 伊藤マンショ |      | 西６         |
	//|   17 | 足利義昭     |      | [1]->X       |
	//|   18 | 大友宗麟     |      | 西４         |
	//+------+--------------+------+--------------+

}
