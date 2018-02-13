package infra

import (
	"github.com/tamura2004/nobugo/adapter"
	"github.com/tamura2004/nobugo/adapter/driver"
	"github.com/tamura2004/nobugo/domain/entity"
	"github.com/tamura2004/nobugo/usecase/port"
)

func Init() {
	entity.Samurai = LoadDeck("config/samurai.toml")
	entity.Castle = LoadDeck("config/castle.toml")

	port.Input = UI{}

	driver.Table = Table{}

	adapter.Init()
}
