package player

type Players struct {
	Num    int
	Active int
	Deck   []Player
}

func Deck() (p Players) {
	num := ui.InputNumber("プレイヤー人数？", 3, 6)

	p.Num = num
	p.Deck = make([]Player, num)
	for i := 0; i < num; i++ {

		p.Deck[i] = Player{}
	}

}
