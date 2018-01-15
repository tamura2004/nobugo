package ability

type ChangeDice struct {
	from int
	to   int
}

func NewChangeDice(x, y int) *ChangeDice {
	return &ChangeDice{
		from: x,
		to:   y,
	}
}
