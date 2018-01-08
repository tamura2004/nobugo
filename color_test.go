package main

import "testing"

func TestColorOwner(t *testing.T) {
	cases := []struct {
		c Color
		o Owner
	}{
		{BLACK, PLAYER1},
		{WHITE, PLAYER2},
		{RED, PLAYER3},
		{BLUE, PLAYER4},
		{GREEN, PLAYER5},
		{YELLOW, PLAYER6},
	}

	for _, c := range cases {
		color, owner := c.c, c.o
		if got := color.Owner(); got != owner {
			t.Errorf("got %v want %v", got, owner)
		}
	}

}
