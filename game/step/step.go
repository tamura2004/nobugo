package step

import "strconv"

type Step int

const (
	PREPARE Step = iota
	MARCH
	EMPLOY
	BATTLE
	CHECK
)

const name = "準備行軍調略合戦終了"

var index = [...]uint8{0, 6, 12, 18, 24, 30}

func (i Step) String() string {
	if i < 0 || i > Step(len(index)-1) {
		return "Step(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return name[index[i]:index[i+1]] + "フェイズ"
}
