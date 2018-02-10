package domain

type SamuraiDeckInterface interface {
	Draw() Samurai
	DrawN(int) SamuraiDeckInterface
	Shuffle()
	Bottom(Samurai)
	BottomN(SamuraiDeckInterface)
}

var SamuraiDeck SamuraiDeckInterface

func InitSamuraiDeck(sd SamuraiDeckInterface) {
	SamuraiDeck = sd
}
