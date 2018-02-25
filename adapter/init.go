package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

func Init() {
	domain.PartyFactory = partyFactory(NewParty)
	domain.BoardFactory = boardFactory(NewBoard)
}
