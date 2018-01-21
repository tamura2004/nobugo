package player

import (
	"fmt"
	"github.com/tamura2004/nobugo/game/board"
	. "github.com/tamura2004/nobugo/player/color"
	"github.com/tamura2004/nobugo/player/pool"

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

func max(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func (p *Player) Prepare() {
	i := len(p.Samurais)
	j := len(p.Castles)
	p.Pool.Num = 3 + max(i, j)
}

func (p *Player) March() {
	fmt.Printf("行軍フェイズ：%s\n", p.Name())

	if p.Pool.Num <= 0 {
		return
	}
	p.Pool.Roll()
	Print()

	p.Pool.Do(func(dice, num int) {
		board.Bid(p.Color, dice, num)
		p.Pool.Num -= num
	})

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
