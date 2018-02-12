package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type PartyConverter struct {
	*domain.Party
}

func NewPartyConverter(p *domain.Party) *PartyConverter {
	return &PartyConverter{
		Party: p,
	}
}

func (p *PartyConverter) ToDic() Dic {
	dic := Dic{}

	p.Each(func(p *domain.Player) {
		pc := PlayerConverter{}
		dic = append(dic, pc.Map(p))
	})

	return dic
}
