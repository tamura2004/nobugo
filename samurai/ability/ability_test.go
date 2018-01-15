package ability

import "testing"

func TestNew(t *testing.T) {
	cases := []struct {
		got  *Ability
		want string
	}{
		{New(CHANGE_DICE, 1, 2, 6), "[1]/[2]->[6]"},
		{New(CHANGE_SAMURAI, 1, 1, 1), "交換：武将"},
		{New(CHANGE_AREA, 1, 1, 1), "西1"},
		{New(ADD_GOLD, 1, 1, 1), "兵糧＋1Ｄ"},
		{New(ADD_BATTLE_DICE, 1, 1, 1), "合戦＋1Ｄ"},
		{New(DELETE_DICE, 1, 1, 1), "1 -> X"},
	}

	for _, c := range cases {
		if c.got.String() != c.want {
			t.Errorf("wrong new ability, got %q want %q", c.got, c.want)
		}
	}

}
