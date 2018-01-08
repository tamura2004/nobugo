package main

import "testing"

func TestColorString(t *testing.T) {
	cases := []struct {
		c Color
		s string
	}{
		{BLACK, "黒"},
		{WHITE, "白"},
		{RED, "赤"},
		{BLUE, "青"},
		{GREEN, "緑"},
		{YELLOW, "黄"},
	}

	for _, c := range cases {
		color, want := c.c, c.s
		if got := color.String(); got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
