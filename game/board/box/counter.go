package box

import (
	. "github.com/tamura2004/nobugo/player/color"
)

type counter map[Color]int

func (c *counter) max() int {
	var x int
	for _, v := range *c {
		if v > x {
			x = v
		}
	}
	return x
}
