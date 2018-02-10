package domain

type DeckInterface interface {
	Draw() Card
	DrawN(int) Deck
	Shuffle()
	Bottom(Card) Deck
	BottomN(Deck) Deck
}

var deck struct {
	samurai Deck
	jikisan Deck
	castle  Deck
}

func InitDeck(s, j, c DeckInterface) {
	deck.samurai = s
	deck.jikisan = j
	deck.castle = c
}
