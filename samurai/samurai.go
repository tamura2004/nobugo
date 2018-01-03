package samurai

import (
	"github.com/tamura2004/nobugo/deck"
)

type Samurai struct {
	number  int
	name    string
	jikisan bool
	ability string
}

// impliment Card interface
func (s *Samurai) Name() string {
	return s.name
}

func NewJikisanDeck() deck.Deck {
	deck := deck.Deck{}
	for i, name := range samuraiNames() {
		if i < 6 {
			samurai := &Samurai{
				number: i,
				name:   name,
			}
			deck = append(deck, samurai)
		}
	}
	deck.Shuffle()
	return deck
}

func samuraiNames() []string {
	return []string{
		"木下藤吉郎",
		"明智光秀",
		"柴田勝家",
		"前田利家",
		"滝川一益",
		"弥助",
		"毛利元就",
		"松永久秀",
		"武田信玄",
		"徳川家康",
		"今川義元",
		"北条氏康",
		"上杉謙信",
		"伊達政宗",
		"九鬼義隆",
		"島津義弘",
		"伊藤マンショ",
		"足利義昭",
		"大友宗麟",
	}
}
