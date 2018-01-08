package main

// DeckはCardを順に積み重ねたもの、所謂山札や手札、場札を表す
type Deck []Card

/*
// draw(n)は山札からn枚カードを引く。十分なカードがないとエラー。
// 副作用で山札を変更する。
func (d *Deck) Draw(n int) (*Deck, error) {
	x := *d
	if len(x) < n {
		return nil, fmt.Errorf("not enough card got %d want %d", len(x), n)
	}
	cards, x := x[:n], x[n:]
	*d = x
	return &cards, nil
}

// DrawOne()は山札からカードを一枚引く
func (d *Deck) DrawOne() (*Card, error) {
	x, err := d.Draw(1)
	if err != nil {
		return nil, err
	}

	return &(*x)[0], nil
}

// Shuffle()は山札をシャッフルする。
func (d *Deck) Shuffle() {
	x := *d
	n := len(x)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
	*d = x
}

func (d Deck) Header() []string {
	if len(d) == 0 {
		return nil
	}
	return d[0].Header()
}

func (d Deck) Data() [][]string {
	x := make([][]string, len(d))
	for i, card := range d {
		x[i] = card.Values()
	}
	return x
}
*/
