package adapter

import (
	"fmt"
	"github.com/tamura2004/nobugo/domain"
	"log"
)

type PlayerConverter struct {
	domain.Player
}

func (p *PlayerConverter) Map() map[string][]string {
	return map[string][]string{
		"色":   p.Color,
		"ダイス": p.Pool,
		"武将":  p.Samurai,
		"城":   p.Castle,
	}
}

func (p *PlayerConverter) Color() []string {
	return []string{p.Color.String()}
}

func (p *PlayerConverter) Pool() []string {
	str := ""
	for _, dice := range p.Pool {
		str += fmt.Sprintf("[%d]")
	}
	return str
}

func (p *PlayerConverter) Samurai() []string {
	arr := []string{}
	p.Samurai.Each(func(c domain.Card) {
		arr = append(arr, c.Name)
	})
	return arr
}

func (p *PlayerConverter) Castle() []string {
	arr := []string{}
	p.Castle.Each(func(c domain.Card) {
		arr = append(arr, c.Name)
	})
	return arr
}
