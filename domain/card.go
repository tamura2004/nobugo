package domain

import (
	"fmt"
)

type Card struct {
	Type    Type
	Area    Area
	Country string
	Name    string
	Ability Ability
	Power   []int
	Done    bool
}

func EmptyCard() Card {
	return Card{
		Type: EMPTY,
	}
}

func HonojiCard() Card {
	return Card{
		Type: HONOJI,
	}
}

func (c *Card) CantChangeDice() bool {
	return c.Done || c.Ability != CHANGE_DICE
}

func (c *Card) CanChangeDice() bool {
	return !c.CantChangeDice()
}

func (c *Card) ChangeFrom() []int {
	return []int{c.Power[0], c.Power[1]}
}

func (c *Card) ChangeTo() int {
	return c.Power[2]
}

func (c *Card) ChangeDiceActions(p *Player) []Action {
	actions := []Action{}
	if c.CantChangeDice() {
		return actions
	}
	for _, n := range c.ChangeFrom() {
		if p.Pool.Include(n) {
			actions = append(actions, c.ChangeDiceAction(p, n, c.ChangeTo()))
		}
	}
	return actions
}

func (c *Card) ChangeDiceAction(p *Player, from, to int) Action {
	return Action{
		Msg: fmt.Sprintf("%s:[%d]->[%d]", c.Name, from, to),
		Do: func() {
			p.Pool.Replace(from, to)
			c.Done = true
		},
	}
}
