package adapter

import (
	"fmt"
	"github.com/tamura2004/nobugo/domain"
)

type Player struct {
	*domain.Player
}

func (p *Player) Map() map[string][]string {
	return map[string][]string{
		"色":   p.Color(),
		"ダイス": p.Pool(),
		"武将":  p.Samurai(),
		"城":   p.Castle(),
	}
}

func (p *Player) Color() []string {
	return []string{p.Player.Color.String()}
}

func (p *Player) Pool() []string {
	str := ""
	for _, dice := range p.Player.Pool {
		str += fmt.Sprintf("[%d]", dice.Num)
	}
	return []string{str}
}

func (p *Player) Samurai() []string {
	arr := []string{}
	p.Player.Samurai.Each(func(c domain.Card) {
		arr = append(arr, c.Name)
	})
	return arr
}

func (p *Player) Castle() []string {
	arr := []string{}
	p.Player.Castle.Each(func(c domain.Card) {
		arr = append(arr, c.Name)
	})
	return arr
}
