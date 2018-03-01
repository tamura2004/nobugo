package domain_test

import (
	"github.com/kr/pretty"
	"github.com/tamura2004/nobugo/domain"
)

func ExamplePartyGetPlayerByColor() {
	p := domain.NewParty(6)
	pl := p.GetPlayerByColor(domain.YELLOW)
	pretty.Print(pl)
	// Output:
	// &domain.Player{
	//     Color: 5,
	//     Pool:  &domain.Pool{
	//     },
	//     Samurai: (*domain.Deck)(nil),
	//     Castle:  (*domain.Deck)(nil),
	// }
}
