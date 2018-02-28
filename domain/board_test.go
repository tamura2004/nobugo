package domain_test

import (
	"github.com/kr/pretty"
	"github.com/tamura2004/nobugo/domain"
)

func ExampleBoardPool() {
	b := domain.NewBoard()
	pretty.Print(b.Pool())
	// Output:
	// []*domain.Pool{
	//     &domain.Pool{
	//     },
	//     &domain.Pool{
	//     },
	//     &domain.Pool{
	//     },
	//     &domain.Pool{
	//     },
	//     &domain.Pool{
	//     },
	//     &domain.Pool{
	//     },
	// }
}

func ExampleBoardDeck() {
	b := domain.NewBoard()
	pretty.Print(b.Deck())
	// Output:
	// []*domain.Deck{
	//     &domain.Deck{},
	//     &domain.Deck{},
	//     &domain.Deck{},
	//     &domain.Deck{},
	//     &domain.Deck{},
	//     &domain.Deck{},
	// }
}
