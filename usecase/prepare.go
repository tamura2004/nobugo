package usecase

import (
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/domain/entity"
)

func prepare() bool {
	entity.Party.Each(func(p *domain.Player) {
		p.Pool = domain.NewPool(p.Color, 3)
	})

	step.Next()
	return true
}
