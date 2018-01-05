package deck

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Card interface {
	Header() []string
	Values() []string
	String() string
}

type Deck []Card

func (d *Deck) Names() string {
	w := make([]string, len(*d))
	for i, card := range *d {
		w[i] = card.String()
	}
	return strings.Join(w, ";")
}

func (d *Deck) PrintTable() {
	if len(*d) == 0 {
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader((*d)[0].Header())
	for _, card := range *d {
		table.Append(card.Values())
	}
	table.Render()
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	deck := *d
	n := len(deck)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	*d = deck
}

func (d *Deck) Draw() Card {
	card, deck := (*d)[0], (*d)[1:]
	*d = deck
	return card
}

func (d *Deck) Bottom(c Card) {
	deck := append(*d, c)
	*d = deck
}

func (d *Deck) DrawIf(filter func(c Card) bool) Card {
	deck := *d
	for i, card := range deck {
		if filter(card) {
			c := card
			deck := append(deck[:i], deck[i+1:]...)
			*d = deck
			return c
		}
	}
	return nil
}
