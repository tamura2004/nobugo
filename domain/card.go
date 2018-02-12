package domain

type Card struct {
	Type    Type
	Area    Area
	Country string
	Name    string
	Ability Ability
	Power   []int
}

func EmptyCard() Card {
	return Card{
		Type: EMPTY,
	}
}

func HonojiCard() Card {
	return Card{
		Type: HONOJI,
	}
}
