package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

type Party struct {
	domain.Party
	Printer
}

// プレイヤー人数を決める
// 各プレイヤーに武将カードを１枚ずつ配る
func NewParty(n int) *Party {
	return &Party{
		Party: domain.NewParty(n),
	}
}
