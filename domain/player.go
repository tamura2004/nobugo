package domain

type Player struct {
	Color
	Pool
	Samurai Deck
	Castle  Deck
	Done    bool
}

func NewPlayer(n int) Player {
	return Player{
		Color: Color(n),
	}
}

func (p *Player) Recover() {
	p.Done = false
}
