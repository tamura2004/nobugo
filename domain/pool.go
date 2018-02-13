package domain

type Pool []Dice

func (p *Pool) Add(d Dice) {
	*p = append(*p, d)
}

func (p *Pool) Append(o Pool) {
	for i := 0; i < len(o); i++ {
		p.Add(o[i])
	}
}

func (p Pool) Roll() {
	for i := 0; i < len(p); i++ {
		p[i].Roll()
	}
}

func (p Pool) Restore() {
	for i := 0; i < len(p); i++ {
		p[i].Restore()
	}
}

func NewPool(c Color, n int) Pool {
	pool := Pool{}
	for i := 0; i < n; i++ {
		pool = append(pool, NewDice(c))
	}
	return pool
}
