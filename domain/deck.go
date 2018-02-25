package domain

import (
	"math/rand"
)

type Deck struct {
	Card []Card
}

func (d *Deck) Len() int {
	return len(d.Card)
}

func (d *Deck) Get(i int) Card {
	return d.Card[i]
}

func (d *Deck) Draw() Card {
	deck := d.Card
	if len(deck) == 0 {
		return EmptyCard()
	}
	card, deck := deck[0], deck[1:]
	d.Card = deck
	return card
}

func (d *Deck) DrawN(n int) *Deck {
	deck := Deck{}
	for i := 0; i < n; i++ {
		if c := d.Draw(); c.Type != EMPTY {
			deck.Bottom(c)
		}
	}
	return &deck
}

func (d *Deck) ShuffleN(n int) {
	deck := d.Card
	if n > len(deck) {
		n = len(deck)
	}
	for i := 1; i < n-1; i++ {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	d.Card = deck
}

func (d *Deck) Shuffle() {
	d.ShuffleN(len(d.Card))
}

func (d *Deck) Bottom(c Card) {
	d.Card = append(d.Card, c)
}

func (d *Deck) Each(f func(c Card)) {
	for i := 0; i < len(d.Card); i++ {
		f(d.Card[i])
	}
}

func (d *Deck) Names() []string {
	names := []string{}
	d.Each(func(c Card) {
		names = append(names, c.Name)
	})
	return names
}
