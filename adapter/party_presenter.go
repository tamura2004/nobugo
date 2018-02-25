package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type PartyPresenter struct {
	*domain.Party
}

func (p PartyPresenter) Header() []string {
	return []string{"色", "ダイス", "武将", "城"}
}

func (p PartyPresenter) Encode() Dic {
	dic := Dic{}

	p.Party.Each(func(p *domain.Player) {
		pc := Player{p}
		dic = append(dic, pc.Map())
	})

	return dic
}

func (p PartyPresenter) Print() {
	Print(p)
}
