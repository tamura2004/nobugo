package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

func open() {
	// プレイヤー人数を決める
	n := UI.Num(3, 6, "プレイヤー人数を入力して下さい")
	Party = domain.NewParty(n)
	Board = domain.NewBoard()

	// デッキをシャッフル
	Samurai.Shuffle()
	Castle.Shuffle()

	// 各プレイヤーに武将カードを1枚ずつ配る
	Party.Each(func(p *domain.Player) {
		card := Samurai.Draw()
		p.Samurai.Bottom(card)
	})
}
