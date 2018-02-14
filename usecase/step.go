package usecase

import "strconv"

type Step int

const (
	PREPARE Step = iota
	MARCH
	EMPLOY
	BATTLE
	CHECK
)

var actions = map[Step]func(){
	PREPARE: prepare,
	MARCH:   march,
	EMPLOY:  employ,
	BATTLE:  battle,
	CHECK:   check,
}

func (s *Step) next() {
	actions[*s]()
	step = *s
	*s = Step((step + 1) % 5)
}

func (i Step) String() string {
	name := "準備行軍調略合戦終了"
	index := [...]uint8{0, 6, 12, 18, 24, 30}
	if i < 0 || i > Step(len(index)-1) {
		return "Step(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return name[index[i]:index[i+1]] + "フェイズ"
}