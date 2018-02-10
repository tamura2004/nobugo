package domain

import (
	"math/rand"
)

type Dice struct {
	Num int
	Color
}

func (d *Dice) Roll() {
	d.Num = rand.Intn(6) + 1
}

func (d *Dice) Restore() {
	d.Num = 0
}
