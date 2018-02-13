package usecase

import (
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/domain/entity"
	"github.com/tamura2004/nobugo/usecase/port"
)

func open() {
	// プレイヤー人数を決める
	n := port.Input.Num(3, 6, "プレイヤー人数を入力して下さい")
	entity.Party = domain.NewParty(n)

	// デッキをシャッフル
	entity.Samurai.Shuffle()
	entity.Castle.Shuffle()

	// 各プレイヤーに武将カードを1枚ずつ配る
	entity.Party.Each(func(player *domain.Player) {
		card := entity.Samurai.Draw()
		player.Samurai.Bottom(card)
	})
}
