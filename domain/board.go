package domain

type Board [6]Box

func NewBoard() Board {
	board := Board{}
	for i := 0; i < 6; i++ {
		board[i] = NewBox(i)
	}
	return board
}
