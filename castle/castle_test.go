package castle

import (
	"fmt"
	"github.com/tamura2004/nobugo/castle/area"
	"testing"
)

func TestNew(t *testing.T) {
	c := New("羅馬", "法王庁", 1, 1, 1, area.W6)
	if c.Country != "羅馬" {
		t.Errorf("bad country, got %v want %v", c.Country, "羅馬")
	}
}

func ExampleValue() {
	c := New("羅馬", "法王庁", 1, 1, 1, area.W6)
	fmt.Println(c.Value())
	// Output:
	// [法王庁 [1][1][1] 西６/羅馬]
}
