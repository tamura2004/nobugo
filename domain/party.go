package domain

type Party struct {
	Num    int
	Active int
	Player []Player
}

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

func (p *Party) Each(f func(p *Player)) {
	for i := 0; i < p.Num; i++ {
		f(&p.Player[i])
	}
}

func (p *Party) Recover() {
	p.Each(func(pl *Player) {
		pl.Recover()
	})
}

func (p *Party) Done() bool {
	for i := 0; i < p.Num; i++ {
		if !p.Player[i].Done {
			return false
		}
	}
	return true
}

func (p *Party) ActivePlayer() Player {
	return p.Player[p.Active]
}

func (p *Party) Next() {
	if p.Done() {
		return
	}

	i := p.Active

	for {
		i = (i + 1) % p.Num
		if !p.Player[i].Done {
			p.Active = i
		}
	}
}
