package main

type Owner int

const (
	DECK Owner = iota
	PLAYER1
	PLAYER2
	PLAYER3
	PLAYER4
	PLAYER5
	PLAYER6
	ACTION1
	ACTION2
	ACTION3
	ACTION4
	ACTION5
	ACTION6
	IKKI
	SIEGE
)

func (o Owner) Color() Color {
	return Color(int(o))
}
