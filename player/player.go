package player

import (
	"fmt"
	"github.com/tamura2004/nobugo/game/board"
	. "github.com/tamura2004/nobugo/player/color"
	"github.com/tamura2004/nobugo/player/pool"
	"github.com/tamura2004/nobugo/ui"

	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/samurai"
)

type Player struct {
	Color  Color
	Pool   pool.Pool
	Active bool
	samurai.Samurais
	castle.Castles
}

func New(i int) *Player {
	sd, err := samurai.Draw(1)
	if err != nil {
		panic("no samurai card when init player")
	}
	return &Player{
		Color:    Color(i),
		Samurais: sd,
	}
}

func (p *Player) Prepare() {
	i := len(p.Samurais)
	j := len(p.Castles)

	if i < j {
		p.Pool.Num = 3 + i
	} else {
		p.Pool.Num = 3 + j
	}
}

func (p *Player) March() {
	if p.Pool.Num == 0 {
		return
	}

	p.Pool.Roll()
	Print()

	items, vals := p.Pool.Selection()

	var ix int
	if len(items) > 1 {
		ix = ui.SelectNumber("select use dice", items)
	} else {
		ix = 1
	}

	dice, num := vals[ix-1][0], vals[ix-1][1]
	board.Stat().Box[dice].Bid(p.Color, num)
	p.Pool.Num -= num

	Print()
	board.Print()
	p.Pool.Rolled = false
}

func (p *Player) AddSamurai(s samurai.Samurai) {
	p.Samurais = append(p.Samurais, s)
}

func (p *Player) AddCastle(c castle.Castle) {
	p.Castles = append(p.Castles, c)
}

func (p *Player) Name() []string {
	if p.Active {
		return []string{fmt.Sprintf("手番 -> %s", p.Color.String())}
	}
	return []string{p.Color.String()}
}

func (p *Player) Value() map[string][]string {
	return map[string][]string{
		"プレイヤー": p.Name(),
		"ダイス":   p.Pool.Value(),
		"武将":    p.Samurais.Values(),
		"城":     p.Castles.Values(),
	}
}
