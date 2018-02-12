package adapter

import (
	"github.com/tamura2004/nobugo/usecase"
	"log"
)

type PartyPrinter struct {
	*usecase.Party
	PartyConverter // ToDic()
}

func NewPartyPrinter(p *usecase.Party) *PartyPrinter {
	return &PartyPrinter{
		Party:          p,
		PartyConverter: NewPartyConverter(p),
	}
}

func (p *PartyPrinter) Print() {
	if TablePrinter == nil {
		log.Fatal("adapter.TablePrinter is not initialized")
	}

	dic := p.ToDic()
	keys := []string{"色", "ダイス", "武将", "城"}
	table := ConvertDicToTable(dic, keys)
	TablePrinter.Print(table.Data)
}
