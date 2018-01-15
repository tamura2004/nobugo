package samurai_test

import (
	"fmt"
	"github.com/tamura2004/nobugo/samurai"
	"github.com/tamura2004/nobugo/samurai/ability"
)

func ExampleNew() {
	s := samurai.New("家康", 1, 2, 3, ability.CHANGE_DICE)

	fmt.Printf("%v", s)
	// Output:
	// {家康 [1]/[2]->[3]}
}

func ExampleGet() {
	sd := samurai.Deck()
	s := sd.Get(0)

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
