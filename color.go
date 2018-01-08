package main

type Color int

const (
	_ Color = iota
	BLACK
	WHITE
	RED
	BLUE
	GREEN
	YELLOW
)

func (c Color) Owner() Owner {
	return Owner(int(c))
}
