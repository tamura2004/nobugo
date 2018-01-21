package box_test

import (
	"github.com/tamura2004/nobugo/game/board/box"
	. "github.com/tamura2004/nobugo/player/color"

	"fmt"
	"testing"
)

func TestWiner(t *testing.T) {
	var b box.Box
	b = box.New()

	b.Bid(BLUE, 3)
	b.Bid(RED, 6)
	b.Bid(BLUE, 3)

	got := b.Winer()
	want := BLUE
	if got != want {
		t.Errorf("bad winner, got %v want %v", got, want)
	}
}

func ExampleDetail() {
	b := box.New()
	b.Bid(BLUE, 3)
	b.Bid(RED, 5)
	b.Bid(GREEN, 7)

	fmt.Println(b.Value())
	// Output:
	// [青:3 赤:5 緑:7]
}
