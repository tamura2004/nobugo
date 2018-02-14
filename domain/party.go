package domain

type Party struct {
	Num    int
	Active int
	Player []Player
}

type Work func(*Player)

func NewParty(n int) Party {
	party := Party{
		Num:    n,
		Active: 0,
	}
	for i := 0; i < n; i++ {
		party.Player = append(party.Player, NewPlayer(i))
	}
	return party
}

func (p *Party) Each(f Work) {
	for i := 0; i < p.Num; i++ {
		f(&p.Player[i])
	}
}

func (p *Party) EachHasDice(f Work) {
	for pl := p.TurnPlayer(); p.Done(); p.Next() {
		if pl.NoDice() {
			continue
		}
		f(pl)
	}
}

func (p *Party) TurnPlayer() *Player {
	return &p.Player[p.Active]
}

func (p *Party) Done() bool {
	flag := true
	p.Each(func(p *Player) {
		if p.HasDice() {
			flag = false
		}
	})
	return flag
}

func (p *Party) Next() {
	p.Active = (p.Active + 1) % p.Num
}
