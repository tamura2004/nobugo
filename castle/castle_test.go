package castle

import (
	"fmt"
	"github.com/tamura2004/nobugo/castle/area"
	"testing"
)

func TestNew(t *testing.T) {
	c := New("羅馬", "法王庁", 1, 1, 1, area.W6)
	if c.Country != "羅馬" {
		t.Error("bad new castle")
	}
}

func ExampleValue() {
	c := New("羅馬", "法王庁", 1, 1, 1, area.W6)
	fmt.Println(c.Value())
	// Output:
	// [法王庁 [1][1][1] 西６/羅馬]
}

func TestDeck(t *testing.T) {
	cs := Deck()
	got := cs.Get(0).Country
	want := "羅馬"
	if got != want {
		t.Error("bad new castle")
	}
}
