package samurai

import (
	"github.com/tamura2004/nobugo/samurai/ability"
)

type Samurais []Samurai

func Deck() Samurais {
	return Samurais{
		New("弥助", 1, 2, 6, ability.CHANGE_DICE),
		New("木下藤吉郎", 3, 4, 1, ability.CHANGE_DICE),
		New("明智光秀", 5, 6, 1, ability.CHANGE_DICE),
		New("柴田勝家", 3, 4, 5, ability.CHANGE_DICE),
		New("前田利家", 5, 6, 3, ability.CHANGE_DICE),
		New("滝川一益", 1, 2, 4, ability.CHANGE_DICE),
		New("毛利元就", 5, 6, 4, ability.CHANGE_DICE),
		New("徳川家康", 3, 4, 2, ability.CHANGE_DICE),
		New("今川義元", 5, 6, 1, ability.CHANGE_DICE),
		New("北条氏康", 1, 2, 3, ability.CHANGE_DICE),
		New("島津義弘", 3, 4, 6, ability.CHANGE_DICE),
		New("松永久秀", 0, 0, 0, ability.CHANGE_SAMURAI),
		New("武田信玄", 1, 0, 0, ability.ADD_GOLD),
		New("上杉謙信", 2, 0, 0, ability.ADD_BATTLE_DICE),
		New("伊達政宗", 2, 4, 0, ability.CHANGE_AREA),
		New("九鬼義隆", 2, 6, 0, ability.CHANGE_AREA),
		New("伊藤マンショ", 1, 6, 0, ability.CHANGE_AREA),
		New("足利義昭", 1, 0, 0, ability.DELETE_DICE),
		New("大友宗麟", 1, 4, 0, ability.CHANGE_AREA),
	}
}

func (sd Samurais) Values() (a []string) {
	for _, s := range sd {
		a = append(a, s.Name)
	}
	return a
}

func (sd *Samurais) Draw(i int) (y Samurais) {
	x := *sd
	x, y = x[i:], x[:i]
	*sd = x
	return
}
