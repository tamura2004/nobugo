package main

type Area int

const (
	W6 Area = iota
	W5
	W4
	W3
	W2
	W1
	E1
	E2
	E3
	E4
	E5
	E6
)

func (a Area) String() string {
	switch a {
	case W6:
		return "西６"
	case W5:
		return "西５"
	case W4:
		return "西４"
	case W3:
		return "西３"
	case W2:
		return "西２"
	case W1:
		return "西１"
	case E1:
		return "東１"
	case E2:
		return "東２"
	case E3:
		return "東３"
	case E4:
		return "東４"
	case E5:
		return "東５"
	case E6:
		return "東６"
	}
	panic("不明なエリア")
}
