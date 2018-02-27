package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type PoolPresenter domain.Pool

func (p PoolPresenter) Names() []string {
	names := []string{}
	for _, dice := range p {
		names = append(names, DicePresenter(dice).Name())
	}
	return names
}
