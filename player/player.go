package player

import (
	"fmt"

	. "github.com/tamura2004/nobugo/player/color"
	. "github.com/tamura2004/nobugo/player/pool"

	"github.com/tamura2004/nobugo/castle"
	"github.com/tamura2004/nobugo/samurai"
	"github.com/tamura2004/nobugo/table"
	"github.com/tamura2004/nobugo/ui/mock"
)

type Player struct {
	Color    Color
	Pool     Pool
	Active   bool
	Samurais Samurais
	Castles  Castles
}

var num int = func() int {
	ui.MsgBox("信長さんの野望")
	return ui.InputNumber("プレイヤー人数？", 3, 6)
}()

var player Players

func init() {
	player.Num = num
	player.Deck = make([]*Player, num)
	for i := 0; i < num; i++ {
		player.Deck[i] = &Player{
			Color: Color(i),
			Owner: Owner(i + 1),
			Pool:  Pool{},
		}
	}
}

func (p *Player) playerName() []string {
	if player.Deck[player.Active].Owner == p.Owner {
		return []string{fmt.Sprintf("手番 -> %s", p.Color.String())}
	}
	return []string{p.Color.String()}
}

func (p *Player) Value() map[string][]string {
	return map[string][]string{
		"プレイヤー": p.playerName(),
		"ダイス":   {fmt.Sprint(p.Pool.Num)},
		"武将":    p.playerSamurai(),
	}
}

func (ps *Players) Values() []map[string][]string {
	ret := []map[string][]string{}
	for _, p := range player.Deck {
		ret = append(ret, p.Value())
	}
	return ret
}

func PrintPlayer() {
	fmt.Println("\n---- プレイヤー一覧 ----")
	table.Render(&player, []string{"プレイヤー", "ダイス", "武将"})
}
