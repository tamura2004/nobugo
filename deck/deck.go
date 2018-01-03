package deck

import (
	"math/rand"
	"time"
)

type Card interface {
	Name() string
}

type Deck []Card

func (d Deck) Shuffle() Deck {
	rand.Seed(time.Now().UnixNano())
	n := len(d)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}
