package box

import (
	"fmt"
	. "github.com/tamura2004/nobugo/player/color"
	"github.com/tamura2004/nobugo/tool/counter"
	"github.com/tamura2004/nobugo/tool/list"
)

// Box is place where each player bid with a dice to get samurai or castle card.
// put more dice and followed later gives high priority.
type Box struct {
	list.List
	counter.Counter
}

func New() Box {
	return Box{
		List:    list.New(),
		Counter: counter.New(),
	}
}

// bid n votes for No_i target
func (b *Box) Bid(c Color, n int) {
	b.Put(c)
	b.Add(c, n)
}

func (b *Box) Winer() (w Color) {
	for i, n := range b.Counter {
		if n == b.Max() {
			if w, ok := i.(Color); ok {
				return Color(w)
			}
		}
	}
	return Color(-1)
}

func (b *Box) Value() (v []string) {
	for _, e := range b.List {
		if c, ok := e.(Color); ok {
			v = append(v, fmt.Sprintf("%s:%d", c, b.Counter[c]))
		}
	}
	return
}
