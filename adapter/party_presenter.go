package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type PartyPresenter domain.Party

func (p PartyPresenter) Header() []string {
	return []string{"色", "ダイス", "武将", "城"}
}

func (p PartyPresenter) Encode() Dic {
	dic := Dic{}

	for _, pl := range p.Player {
		pc := PlayerPresenter(pl)
		dic = append(dic, pc.Map())
	}

	return dic
}

func PartyPrint(p domain.Party) {
	Print(PartyPresenter(p))
}
