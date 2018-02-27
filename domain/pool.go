package domain

type Pool []Dice

func NewPool(c Color, n int) *Pool {
	pool := &Pool{}
	for i := 0; i < n; i++ {
		*pool = append(*pool, NewDice(c))
	}
	return pool
}

func NewEmptyPool() *Pool {
	return &Pool{}
}

func (p *Pool) ToMap() map[int]int {
	m := make(map[int]int)
	for _, d := range *p {
		m[d.Num]++
	}
	return m
}

func (p *Pool) Add(d Dice) {
	*p = append(*p, d)
}

func (p *Pool) Move(dst *Pool, n int) {
	src := Pool{}
	for _, dice := range *p {
		if dice.Num == n {
			dst.Add(dice)
		} else {
			src.Add(dice)
		}
	}
	*p = src
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
			return
		}
	}
}
