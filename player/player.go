package player

import (
	"fmt"
	"strconv"

	. "github.com/tamura2004/nobugo/player/color"
	"github.com/tamura2004/nobugo/player/pool"

	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/samurai"
)

type Player struct {
	Color    Color
	Pool     pool.Pool
	Active   bool
	Samurais samurai.Samurais
	Castles  castle.Castles
}

func New(i int) Player {
	return Player{
		Color: Color(i),
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

func (p *Player) Name() []string {
	if p.Active {
		return []string{fmt.Sprintf("手番 -> %s", p.Color.String())}
	}
	return []string{p.Color.String()}
}

func (p *Player) Value() map[string][]string {
	return map[string][]string{
		"プレイヤー": p.Name(),
		"ダイス":   {strconv.Itoa(p.Pool.Num)},
		"武将":    p.Samurais.Values(),
	}
}
