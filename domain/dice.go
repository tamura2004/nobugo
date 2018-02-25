package domain

import (
	"fmt"
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

func NewDice(c Color) Dice {
	return Dice{
		Num:   1,
		Color: c,
	}
}

func (d *Dice) Name() string {
	return fmt.Sprintf("[%s%d]", d.Color, d.Num)
}
