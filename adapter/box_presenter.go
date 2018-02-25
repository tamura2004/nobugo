package adapter

import (
	"github.com/tamura2004/nobugo/domain"
	"strconv"
)

type BoxPresenter struct {
	*domain.Box
}

func (p BoxPresenter) Map() map[string][]string {
	b := p.Box
	m := make(map[string][]string)

	m["番号"] = []string{strconv.Itoa(b.Num + 1)}
	m["名称"] = []string{domain.BoxName(b.Num)}
	m["カード"] = b.Deck.Names()
	m["ダイス"] = b.Pool.Names()

	return m
}
