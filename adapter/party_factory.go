package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type partyFactory func(int) *domain.Party

func (f partyFactory) Create(n int) *domain.Party {
	return f(n)
}

func NewParty(n int) *domain.Party {
	p := domain.NewParty(n)

	for i := 0; i < n; i++ {
		p.Player[i] = domain.NewPlayer(i)
	}

	p.Printer = PartyPresenter{p}

	return p
}
