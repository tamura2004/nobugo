package domain_test

import (
	"github.com/tamura2004/nobugo/domain"
	"testing"
)

func TestAbilityString(t *testing.T) {
	cases := []struct {
		abi domain.Ability
		str string
	}{
		{domain.CHANGE_DICE, "CHANGE_DICE"},
		{domain.CHANGE_SAMURAI, "CHANGE_SAMURAI"},
		{domain.CHANGE_AREA, "CHANGE_AREA"},
		{domain.ADD_GOLD, "ADD_GOLD"},
		{domain.ADD_BATTLE_DICE, "ADD_BATTLE_DICE"},
		{domain.DELETE_DICE, "DELETE_DICE"},
	}

	for _, c := range cases {
		abi, str := c.abi, c.str
		got := abi.String()
		want := str
		if got != want {
			t.Errorf("wrong ability to string, got %q want %q", got, want)
		}
	}

	for _, c := range cases {
		abi, str := c.abi, c.str
		got := domain.Ability(0)
		got.UnmarshalText([]byte(str))
		want := abi
		if got != want {
			t.Errorf("wrong string to ability , got %q want %q", got, want)
		}
	}

}
