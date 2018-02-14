package domain

type Player struct {
	Color
	Pool
	Samurai Deck
	Castle  Deck
}

type Effect func(*Card)

func NewPlayer(n int) Player {
	return Player{
		Color: Color(n),
	}
}

func (p *Player) DiceBonus() int {
	s := p.Samurai.Len()
	c := p.Castle.Len()
	if s > c {
		return c
	}
	return s
}

func (p *Player) NoDice() bool {
	if len(p.Pool) == 0 {
		return true
	}
	return false
}

func (p *Player) HasDice() bool {
	return !p.NoDice()
}

func (p *Player) EachSamurai(f Effect) {
	deck := p.Samurai
	for i := 0; i < deck.Len(); i++ {
		f(&deck.Card[i])
	}
}
