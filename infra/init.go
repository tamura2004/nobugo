package infra

import (
	"github.com/tamura2004/nobugo/adapter"
	"github.com/tamura2004/nobugo/usecase"
)

func Init() {
	usecase.Samurai = LoadDeck("config/samurai.toml")
	usecase.Castle = LoadDeck("config/castle.toml")

	usecase.UI = UI{}

	adapter.TableDriver = Table{}

	adapter.Init()
}
