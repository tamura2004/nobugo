package domain

type Dice struct {
	Num int
	Color
}

func (d *Dice) Roll() {
	d.Num = Rand.Intn(6) + 1
}

func (d *Dice) Restore() {
	d.Num = 0
}

func NewDice(c Color) Dice {
	return Dice{
		Num:   0,
		Color: c,
	}
}
