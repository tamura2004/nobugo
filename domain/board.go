package domain

type Board [6]Box

func NewBoard() *Board {
	b := Board{}
	for i := 0; i < 6; i++ {
		b[i] = NewBox(i)
	}
	return &b
}

var BoardPrint func(Board)

func (b Board) Print() {
	BoardPrint(b)
}

func (b *Board) GetPool(n int) *Pool {
	return b[n].Pool
}

func (b *Board) Pool() []*Pool {
	p := []*Pool{}
	for i := 0; i < 6; i++ {
		p = append(p, b[i].Pool)
	}
	return p
}

func (b *Board) Deck() []*Deck {
	d := []*Deck{}
	for i := 0; i < 6; i++ {
		d = append(d, b[i].Deck)
	}
	return d
}
