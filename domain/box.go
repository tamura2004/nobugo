package domain

type Box struct {
	Name string
	Num  int
	Deck
	Pool
}

func NewBox(n int) Box {
	return Box{
		Name: BoxName(n),
		Num:  n,
	}
}

func BoxName(n int) string {
	if n <= 2 {
		return "調略"
	}
	return "合戦"
}
