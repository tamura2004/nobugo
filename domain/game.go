package domain

type Game struct {
	Turn int
	Step
}

func NewGame() Game {
	return Game{
		Turn: 1,
		Step: PREPARE,
	}
}
