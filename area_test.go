package main

import "testing"

func TestArea(t *testing.T) {
	cases := []struct {
		a Area
		s string
	}{
		{W6, "西６"},
		{W5, "西５"},
		{W4, "西４"},
		{W3, "西３"},
		{W2, "西２"},
		{W1, "西１"},
		{E1, "東１"},
		{E2, "東２"},
		{E3, "東３"},
		{E4, "東４"},
		{E5, "東５"},
		{E6, "東６"},
	}

	for _, c := range cases {
		got, want := c.a.String(), c.s
		if got != want {
			t.Errorf("wrong string, got %q want %q", got, want)
		}
	}
}
