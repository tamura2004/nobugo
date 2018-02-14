package usecase

import (
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/domain/entity"
)

func prepare() {
	// 各プレイヤーはダイスを３個得る
	// 城と武将の組一つ事に追加１ダイスを得る
	entity.Party.Each(func(p *domain.Player) {
		p.Pool = domain.NewPool(p.Color, 3+p.DiceBonus())
	})
}
