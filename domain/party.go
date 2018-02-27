package domain

type Party struct {
	Num    int
	Active int
	Player []Player
}

func NewParty(n int) *Party {
	p := &Party{
		Num:    n,
		Active: 0,
		Player: make([]Player, n),
	}
	for i := 0; i < n; i++ {
		p.Player[i] = NewPlayer(i)
	}
	return p
}

type Work func(*Player)

func (p *Party) Each(f Work) {
	for i := 0; i < p.Num; i++ {
		f(&p.Player[i])
	}
}

func (p *Party) EachHasDice(f Work) {
	for {
		if p.Done() {
			return
		}
		player := p.TurnPlayer()
		if player.NoDice() {
			continue
		}
		f(player)
		p.Next()
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
