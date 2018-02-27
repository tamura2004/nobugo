package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type PlayerPresenter domain.Player

func (p PlayerPresenter) Map() map[string][]string {
	return map[string][]string{
		"色":   {p.Color.String()},
		"ダイス": PoolPresenter(*p.Pool).Names(),
		"武将":  DeckPresenter(p.Samurai).Names(),
		"城":   DeckPresenter(p.Castle).Names(),
	}
}
