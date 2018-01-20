package color

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

const name = "黒白赤青緑黄"

var index = [...]uint8{0, 3, 6, 9, 12, 15, 18}

func (i Color) String() string {
	if i < 0 || i >= Color(len(index)-1) {
		return "Color(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return name[index[i]:index[i+1]]
}

func (c Color) Title() string {
	return c.String()
}

func (c Color) Detail() []string {
	return []string{
		c.String(),
	}
}
