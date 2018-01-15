package counter_test

import (
	"github.com/tamura2004/nobugo/tool/counter"
	"testing"
)

func TestCounterMax(t *testing.T) {
	var c counter.Counter
	c = counter.New()

	c.Add(3, 6)
	c.Add(5, 4)
	c.Add(7, 2)

	got := c.Max()
	want := 6

	if got != want {
		t.Errorf("bad counter max(), got %v want %v", got, want)
	}
}
