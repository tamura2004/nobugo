package domain

type Board [6]Box

func NewBoard() *Board {
	b := Board{}
	for i := 0; i < 6; i++ {
		b[i] = NewBox(i)
	}
	return &b
}

func (b *Board) GetPool(n int) *Pool {
	return b[n].Pool
}
