package infra

import (
	"github.com/BurntSushi/toml"
	"github.com/tamura2004/nobugo/domain"
	"log"
)

func LoadDeck(file string) domain.Deck {
	var deck domain.Deck

	_, err := toml.DecodeFile(file, &deck)
	if err != nil {
		log.Fatal(err)
	}

	return deck
}
