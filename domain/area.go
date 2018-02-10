package domain

import "strconv"

type Area int

const (
	W6 Area = iota + 1
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
	if a > 6 {
		return "東" + strconv.FormatInt(int64(int(a)-6), 10)
	} else {
		return "西" + strconv.FormatInt(int64(7-int(a)), 10)
	}
}
