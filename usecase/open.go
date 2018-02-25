package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

func open() {
	// プレイヤー人数を決める
	n := UI.Num(3, 6, "プレイヤー人数を入力して下さい")
	Party = domain.PartyFactory.Create(n)
	Board = domain.BoardFactory.Create()

	// デッキをシャッフル
	Samurai.Shuffle()
	Castle.Shuffle()

	// 各プレイヤーに武将カードを1枚ずつ配る
	Party.Each(func(player *domain.Player) {
		card := Samurai.Draw()
		player.Samurai.Bottom(card)
	})
}
