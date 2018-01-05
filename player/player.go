package player

//import "strings"
import (
	"fmt"
	//	"math/rand"
	"github.com/tamura2004/nobugo/deck"
)

type Current bool

func (c Current) String() string {
	if c {
		return "レ"
	} else {
		return " "
	}
}

type Player struct {
	Current
	Name    string
	Samurai deck.Deck
	Castle  deck.Deck
	Pool    []int
}

func (p Player) String() string {
	return p.Name
}

func (p Player) Header() []string {
	return []string{"手番", "名前", "武将", "城", "ダイス"}
}

func (p Player) Values() []string {
	return []string{
		p.Current.String(),
		p.Name,
		p.Samurai.Names(),
		p.Castle.Names(),
		fmt.Sprintf("%v", p.Pool),
	}
}
