package usecase

import (
	"fmt"
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/domain/entity"
	"github.com/tamura2004/nobugo/usecase/port"
)

func march() {
	// 開始プレイヤーからスタート
	// 出目を一つ選んで投票ボックスに置く
	// 次のプレイヤーへ
	entity.Party.EachHasDice(func(p *domain.Player) {
		// ダイスを残しているプレイヤーは全てのダイスを振る
		p.Roll()

		// 武将能力を好きなだけ使用する
		action := []domain.Action{}
		p.EachSamurai(func(c *domain.Card) {
			if c.Ability != domain.CHANGE_DICE {
				return
			}
			pow := c.Power
			for _, n := range pow[:1] {
				if p.Pool.Include(n) {
					a := domain.Action{
						Msg: fmt.Sprintf("[%d]->[%d]", n, pow[2]),
						Do: func(x, y int) func() {
							return func() {
								p.Pool.Replace(x, y)
							}
						}(n, pow[2]),
					}
					action = append(action, a)
				}
			}
		})
		port.Input.Select(action)
	})
}
