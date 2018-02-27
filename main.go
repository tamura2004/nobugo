package main

import (
	"github.com/tamura2004/nobugo/adapter"
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/infra"
	"github.com/tamura2004/nobugo/usecase"
)

func main() {
	// 乱数生成の初期化
	domain.Rand = infra.Rand{}

	// デッキの初期化
	usecase.Samurai = infra.LoadDeck("config/samurai.toml")
	usecase.Castle = infra.LoadDeck("config/castle.toml")

	// UIとビューの初期化
	usecase.UI = infra.UI{}
	adapter.TableDriver = infra.Table{}

	// 出力portの初期化
	usecase.PartyPrint = adapter.PartyPrint
	usecase.BoardPrint = adapter.BoardPrint

	// Run
	usecase.Run()
}
