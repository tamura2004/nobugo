package ability

import (
	"fmt"
)

type Ability struct {
	Type int
	p    [3]int
}

func New(t, x, y, z int) *Ability {
	return &Ability{
		Type: t,
		p:    [3]int{x, y, z},
	}
}

func (a Ability) String() string {
	x, y, z := a.p[0], a.p[1], a.p[2]
	switch a.Type {
	case CHANGE_DICE:
		return fmt.Sprintf("[%d]/[%d]->[%d]", x, y, z)
	case CHANGE_SAMURAI:
		return "交換：武将"
	case CHANGE_AREA:
		if x == 1 {
			return fmt.Sprintf("西%d", y)
		} else {
			return fmt.Sprintf("東%d", y)
		}
	case ADD_GOLD:
		return fmt.Sprintf("兵糧＋%dＤ", x)
	case ADD_BATTLE_DICE:
		return fmt.Sprintf("合戦＋%dＤ", x)
	case DELETE_DICE:
		return fmt.Sprintf("%d -> X", x)
	}
	panic("不明な能力タイプ")
}

const (
	CHANGE_DICE = iota
	CHANGE_SAMURAI
	CHANGE_AREA
	ADD_GOLD
	ADD_BATTLE_DICE
	DELETE_DICE
)
