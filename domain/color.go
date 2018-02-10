package domain

import "strconv"

type Color int

const (
	BLACK Color = iota
	WHITE
	RED
	BLUE
	GREEN
	YELLOW
)

func (i Color) String() string {
	name := "黒白赤青緑黄"
	index := [...]uint8{0, 3, 6, 9, 12, 15, 18}
	if i < 0 || i >= Color(len(index)-1) {
		return "Color(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return name[index[i]:index[i+1]]
}
