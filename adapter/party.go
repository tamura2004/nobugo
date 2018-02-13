package adapter

import (
	"github.com/tamura2004/nobugo/adapter/driver"
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/domain/entity"
	"log"
)

type Party struct{}

func (p Party) Print() {
	if driver.Table == nil {
		log.Fatal("driver.Table is not initialized")
	}

	dic := p.ToDic()
	keys := []string{"色", "ダイス", "武将", "城"}
	table := ConvertDicToTable(dic, keys)
	driver.Table.Print(table.Data)
}

func (p Party) ToDic() Dic {
	dic := Dic{}

	entity.Party.Each(func(p *domain.Player) {
		pc := Player{p}
		dic = append(dic, pc.Map())
	})

	return dic
}
