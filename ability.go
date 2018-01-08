package main

import (
	"fmt"
)

type Ability struct {
	Type AbilityType
	p    [3]int
}

func NewAbility(t AbilityType, x, y, z int) Ability {
	return Ability{
		Type: t,
		p:    [3]int{x, y, z},
	}
}

func (a Ability) String() string {
	x, y, z := a.p[0], a.p[1], a.p[2]
	switch a.Type {
	case DEFENCE:
		if z != 0 {
			return fmt.Sprintf("[%d][%d][%d]", x, y, z)
		} else if y != 0 {
			return fmt.Sprintf("[%d][%d]", x, y)
		} else {
			return fmt.Sprintf("[%d]", x)
		}
	case CHANGE_DICE:
		return fmt.Sprintf("[%d/%d -> %d]", x, y, z)
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

type AbilityType int

const (
	DEFENCE AbilityType = iota
	CHANGE_DICE
	CHANGE_SAMURAI
	CHANGE_AREA
	ADD_GOLD
	ADD_BATTLE_DICE
	DELETE_DICE
)

func (a AbilityType) String() string {
	switch a {
	case DEFENCE:
		return "防御"
	case CHANGE_DICE:
		return "変更：ダイス"
	case CHANGE_AREA:
		return "変更：地域"
	case ADD_GOLD:
		return "追加：兵糧"
	case ADD_BATTLE_DICE:
		return "追加：攻撃"
	case DELETE_DICE:
		return "削除：ダイス"
	}
	panic("不明な能力タイプ")
}
