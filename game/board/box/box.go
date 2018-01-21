package box

import (
	"fmt"
	. "github.com/tamura2004/nobugo/player/color"
)

// Box is place where each player bid with a dice to get samurai or castle card.
// put more dice and followed later gives high priority.
type Box struct {
	list
	counter
}

func New() Box {
	return Box{
		list:    list{},
		counter: make(counter),
	}
}

// bid n votes for No_i target
func (b *Box) Bid(c Color, n int) {
	b.list = b.list.put(c)
	b.counter[c] += n
}

func (b *Box) Winer() (w Color) {
	max := b.counter.max()
	b.list.reverse()
	for _, color := range b.list {
		if max == b.counter[color] {
			return color
		}
	}
	return Color(-1)
}

func (b *Box) Value() (v []string) {
	for _, color := range b.list {
		v = append(v, fmt.Sprintf("%s:%d", color, b.counter[color]))
	}
	return v
}
