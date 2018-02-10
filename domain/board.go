package domain

type Board struct {
	Cols [6]Col
}

type Col struct {
	Box
	Name string
	Card
}
