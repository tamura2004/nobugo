package usecase

import (
	"fmt"
)

func march() {
	fmt.Println("march start")
	// 開始プレイヤーからスタート
	for {
		// 全てのプレイヤーがダイスを置き終わったらフェイズ終了
		if Party.Done() {
			return
		}

		// ダイスの無いプレイヤーは飛ばす
		player := Party.TurnPlayer()
		if player.NoDice() {
			Party.Next()
			continue
		}

		// ダイスを残しているプレイヤーは全てのダイスを振る
		player.Roll()

		// 武将能力を好きなだけ使用する
		UI.Select(player.ChangeDiceActions())
		Board.Print()
		Party.Print()

		// 出目を一つ選んで投票ボックスに置く
		UI.Select(player.PutDiceActions(Board))
		Board.Print()
		Party.Print()

		// 次のプレイヤーへ
		Party.Next()
	}
}
