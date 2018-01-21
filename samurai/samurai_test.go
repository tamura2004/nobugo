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

func ExampleValue() {
	s := samurai.New("家康", 1, 2, 3, ability.CHANGE_DICE)

	fmt.Printf("%v", s.Value())
	// Output:
	// [家康 [1]/[2]->[3]]
}

func TestDraw(t *testing.T) {
	samurai.InitDeck()
	dr, _ := samurai.Draw(3)

	got := len(dr)
	want := 3

	if got != want {
		t.Errorf("bad draw %v", dr)
	}
}

func TestDrawError(t *testing.T) {
	samurai.InitDeck()
	_, got := samurai.Draw(300)

	want := "Samurai deck is empty"

	if got.Error() != want {
		t.Errorf("bad draw %v", got)
	}
}
