package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

func prepare() {
	// 各プレイヤーはダイスを３個得る
	// 城と武将の組一つ事に追加１ダイスを得る
	Party.Each(func(p *domain.Player) {
		p.Pool = domain.NewPool(p.Color, 3+p.DiceBonus())
	})

	// 投票カード１，２に武将を2枚ずつ置く
	for i := 0; i < 2; i++ {
		Board.Box[i].Deck.Bottom(Samurai.Draw())
		Board.Box[i].Deck.Bottom(Samurai.Draw())
	}

	// 投票カード３～６に城カードを１枚ずつ置く
	for i := 2; i < 6; i++ {
		Board.Box[i].Deck.Bottom(Castle.Draw())
	}
}
