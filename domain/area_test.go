package domain_test

import (
	"github.com/tamura2004/nobugo/domain"
	"testing"
)

func TestArea(t *testing.T) {
	cases := []struct {
		a domain.Area
		s string
	}{
		{domain.W6, "西6"},
		{domain.W5, "西5"},
		{domain.W4, "西4"},
		{domain.W3, "西3"},
		{domain.W2, "西2"},
		{domain.W1, "西1"},
		{domain.E1, "東1"},
		{domain.E2, "東2"},
		{domain.E3, "東3"},
		{domain.E4, "東4"},
		{domain.E5, "東5"},
		{domain.E6, "東6"},
	}

	for _, c := range cases {
		got, want := c.a.String(), c.s
		if got != want {
			t.Errorf("wrong string, got %q want %q", got, want)
		}
	}
}
