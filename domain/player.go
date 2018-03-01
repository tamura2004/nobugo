package domain

import (
	"fmt"
)

type Player struct {
	Color
	*Pool
	Samurai *Deck
	Castle  *Deck
}

type Effect func(*Card)

func NewPlayer(n int) Player {
	return Player{
		Color:   Color(n),
		Pool:    NewEmptyPool(),
		Samurai: NewDeck(),
		Castle:  NewDeck(),
	}
}

func (p *Player) DiceBonus() int {
	s := p.Samurai.Len()
	c := p.Castle.Len()
	if s > c {
		return c
	}
	return s
}

func (p *Player) NoDice() bool {
	if len(*p.Pool) == 0 {
		return true
	}
	return false
}

func (p *Player) HasDice() bool {
	return !p.NoDice()
}

func (p *Player) EachSamurai(f Effect) {
	deck := p.Samurai
	for i := 0; i < deck.Len(); i++ {
		f(&deck.Card[i])
	}
}

func (p *Player) ChangeDiceActions() []Action {
	actions := []Action{Nop}
	for i := 0; i < p.Samurai.Len(); i++ {
		s := p.Samurai.Get(i)
		actions = append(actions, s.ChangeDiceActions(p)...)
	}
	return actions
}

func (p *Player) PutDiceActions(b *Board) []Action {
	actions := []Action{}
	for n := 1; n <= 6; n++ {
		if p.Pool.Include(n) {
			actions = append(actions, p.PutDiceAction(n, b))
		}
	}
	return actions
}

func (p *Player) PutDiceAction(n int, b *Board) Action {
	return Action{
		Msg: fmt.Sprintf("put dice [%d] to board", n),
		Do: func() {
			p.Pool.Move(b.GetPool(n-1), n)
		},
	}
}
