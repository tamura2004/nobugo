package infra

import (
	"github.com/tamura2004/nobugo/domain/entity"
)

func Init() {
	entity.Samurai = LoadDeck("toml/samurai.toml")
	entity.Castle = LoadDeck("toml/castle.toml")
}
