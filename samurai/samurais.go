package samurai

import (
	"errors"
	"github.com/tamura2004/nobugo/samurai/ability"
	"math/rand"
	"time"
)

type Samurais []Samurai

func init() {
	rand.Seed(time.Now().UnixNano())
}

var deck Samurais

func InitDeck() {
	deck = Samurais{
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
	Shuffle(6)
}

func Draw(num int) (Samurais, error) {
	var y Samurais
	if len(deck) < num {
		return y, errors.New("Samurai deck is empty")
	}
	deck, y = deck[num:], deck[:num]
	return y, nil
}

func DrawOne() (Samurai, error) {
	sd, err := Draw(1)
	if err != nil {
		return Samurai{}, err
	}
	return sd[0], nil
}

// shuffle top num card of deck
func Shuffle(num int) {
	for i := 1; i < num-1; i++ {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func ShuffleAll() {
	Shuffle(len(deck))
}

func (sd Samurais) Values() (a []string) {
	for _, s := range sd {
		a = append(a, s.Name)
	}
	return a
}
