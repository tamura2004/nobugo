package player

import (
	"fmt"
	"github.com/tamura2004/nobugo/table"
)

type Players []Player

func Setup(num int) (ps Players) {
	for i := 0; i < num; i++ {
		ps = append(ps, New(i))
	}
	return ps
}

func (ps Players) Table() (table []map[string][]string) {
	for i := 0; i < len(ps); i++ {
		table = append(table, ps[i].Value())
	}
	return table
}

func (ps Players) Print() {
	fmt.Println("\n---- プレイヤー一覧 ----")
	table.Render(ps, []string{"プレイヤー", "ダイス", "武将"})
}
