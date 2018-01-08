package main

type Step int

const (
	PREPARE Step = iota
	MARCH
	EMPLOY
	BATTLE
	CHECK
)

func (s Step) String() string {
	switch s {
	case PREPARE:
		return "準備フェイズ"
	case MARCH:
		return "行軍フェイズ"
	case EMPLOY:
		return "調略フェイズ"
	case BATTLE:
		return "合戦フェイズ"
	case CHECK:
		return "終了フェイズ"
	}
	return ""
}
