package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type Ability int

const (
	CHANGE_DICE Ability = iota + 1
	CHANGE_SAMURAI
	CHANGE_AREA
	ADD_GOLD
	ADD_BATTLE_DICE
	DELETE_DICE
)

const _Ability_name = "CHANGE_DICECHANGE_SAMURAICHANGE_AREAADD_GOLDADD_BATTLE_DICEDELETE_DICE"

var _Ability_index = [...]uint8{0, 11, 25, 36, 44, 59, 70}

func (i Ability) String() string {
	i -= 1
	if i < 0 || i >= Ability(len(_Ability_index)-1) {
		return "Ability(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Ability_name[_Ability_index[i]:_Ability_index[i+1]]
}

func (a *Ability) UnmarshalText(text []byte) error {
	offset := strings.Index(_Ability_name, string(text))
	for i, ix := range _Ability_index {
		if uint8(offset) == ix {
			*a = Ability(i + 1)
			return nil
		}
	}
	return fmt.Errorf("bad ability string %s", string(text))
}
