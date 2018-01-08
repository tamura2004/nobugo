package main

import "testing"

func TestAbility(t *testing.T) {
	cases := []struct {
		a Ability
		s string
	}{
		{NewAbility(DEFENCE, 3, 2, 1), "[3][2][1]"},
		{NewAbility(DEFENCE, 3, 2, 0), "[3][2]"},
		{NewAbility(DEFENCE, 3, 0, 0), "[3]"},
		{NewAbility(CHANGE_DICE, 1, 2, 6), "[1/2 -> 6]"},
		{NewAbility(CHANGE_SAMURAI, 1, 1, 1), "交換：武将"},
		{NewAbility(CHANGE_AREA, 1, 1, 1), "西1"},
		{NewAbility(ADD_GOLD, 1, 1, 1), "兵糧＋1Ｄ"},
		{NewAbility(ADD_BATTLE_DICE, 1, 1, 1), "合戦＋1Ｄ"},
		{NewAbility(DELETE_DICE, 1, 1, 1), "1 -> X"},
	}

	for _, c := range cases {
		got, want := c.a.String(), c.s
		if got != want {
			t.Errorf("wrong string, got %q want %q", got, want)
		}
	}

}
