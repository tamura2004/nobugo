package box

import (
	"fmt"
	. "github.com/tamura2004/nobugo/player/color"
)

type list []Color

func (ls list) reverse() {
	for i, j := 0, len(ls)-1; i < j; i, j = i+1, j-1 {
		ls[i], ls[j] = ls[j], ls[i]
	}
}

func (ls list) put(c Color) list {
	if ix := ls.find(c); ix != -1 {
		return append(ls.delete(ix), c)
	}
	return append(ls, c)
}

func (ls list) find(c Color) int {
	for i, color := range ls {
		if c == color {
			return i
		}
	}
	return -1
}

func (ls list) delete(ix int) list {
	if ix < 0 || len(ls) <= ix {
		fmt.Printf("out of range, ix = %v\n", ix)
		return ls
	}
	return append(ls[:ix], ls[ix+1:]...)
}
