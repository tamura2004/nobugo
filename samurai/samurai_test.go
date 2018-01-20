package samurai_test

import (
	"fmt"
	"github.com/tamura2004/nobugo/samurai"
	"github.com/tamura2004/nobugo/samurai/ability"
	"testing"
)

func ExampleNew() {
	s := samurai.New("家康", 1, 2, 3, ability.CHANGE_DICE)

	fmt.Printf("%v", s)
	// Output:
	// {家康 [1]/[2]->[3]}
}

func ExampleGet() {
	sd := samurai.Deck()
	s := sd[0]

	fmt.Printf("%v", s)
	// Output:
	// {弥助 [1]/[2]->[6]}
}

func ExampleValue() {
	s := samurai.New("家康", 1, 2, 3, ability.CHANGE_DICE)

	fmt.Printf("%v", s.Value())
	// Output:
	// [家康 [1]/[2]->[3]]
}

func TestDraw(t *testing.T) {
	sd := samurai.Deck()
	n := len(sd)
	dr := sd.Draw(3)

	got := len(dr)
	want := 3

	if got != want {
		t.Errorf("bad draw %v", dr)
	}

	got = len(sd)
	want = n - 3
	if got != want {
		t.Errorf("bad draw %v", dr)
	}

}
