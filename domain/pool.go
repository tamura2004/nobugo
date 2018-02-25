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

func (p Pool) Include(n int) bool {
	for i := 0; i < len(p); i++ {
		if p[i].Num == n {
			return true
		}
	}
	return false
}

func (p Pool) Replace(x, y int) {
	for i := 0; i < len(p); i++ {
		if p[i].Num == x {
			p[i].Num = y
		}
	}
}

func NewPool(c Color, n int) Pool {
	pool := Pool{}
	for i := 0; i < n; i++ {
		pool = append(pool, NewDice(c))
	}
	return pool
}

func (p *Pool) Names() []string {
	names := []string{}
	for _, d := range *p {
		names = append(names, d.Name())
	}
	return names
}
