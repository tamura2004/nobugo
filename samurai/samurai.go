package samurai

import (
	"fmt"

	"github.com/tamura2004/nobugo/deck"
)

type Samurai struct {
	number  int
	name    string
	jikisan bool
	ability Ability
}

func (s Samurai) String() string {
	return s.name
}

func (s Samurai) Header() []string {
	return []string{"番号", "名前", "直臣", "能力"}
}

func (s Samurai) Jikisan() string {
	if s.jikisan {
		return "〇"
	} else {
		return " "
	}
}

func (s Samurai) Values() []string {
	return []string{
		fmt.Sprint(s.number),
		s.name,
		s.Jikisan(),
		s.ability.String(),
	}
}

type Ability interface {
	Type() string
	String() string
}

type Convert struct {
	From int
	To   int
}

func (c Convert) Type() string {
	return "Convert"
}

func (c Convert) String() string {
	return fmt.Sprintf("[%d]/[%d]->[%d]", c.From, c.From+1, c.To)
}

type Special struct {
	text string
}

func (s Special) Type() string {
	return "Special"
}

func (s Special) String() string {
	return s.text
}

func Deck() deck.Deck {
	names := []string{
		"木下藤吉郎",
		"明智光秀",
		"柴田勝家",
		"前田利家",
		"滝川一益",
		"弥助",
		"毛利元就",
		"徳川家康",
		"今川義元",
		"北条氏康",
		"島津義弘",

		"松永久秀",
		"武田信玄",
		"上杉謙信",
		"伊達政宗",
		"九鬼義隆",
		"伊藤マンショ",
		"足利義昭",
		"大友宗麟",
	}

	converts := []Convert{
		Convert{3, 1},
		Convert{5, 2},
		Convert{3, 5},
		Convert{5, 3},
		Convert{1, 4},
		Convert{1, 6},
		Convert{5, 4},
		Convert{3, 2},
		Convert{5, 1},
		Convert{1, 3},
		Convert{1, 5},
	}

	specials := []Special{
		Special{"交換：武将"},
		Special{"城主：Ｄ＋１"},
		Special{"合戦：＋２Ｄ"},
		Special{"東４"},
		Special{"東６"},
		Special{"西６"},
		Special{"[1]->X"},
		Special{"西４"},
	}

	deck := deck.Deck{}

	for i, name := range names {
		samurai := Samurai{
			number:  i,
			name:    name,
			jikisan: i < 6,
		}
		if i < 11 {
			samurai.ability = converts[i]
		} else {
			samurai.ability = specials[i-11]
		}
		deck = append(deck, samurai)
	}

	return deck
}
