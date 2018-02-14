package usecase

import (
	"github.com/tamura2004/nobugo/usecase/port"
)

var step Step
var gameOver bool

func Run() {
	open()
	for !gameOver {
		step.next()
		port.Party.Print()
		port.Board.Print()
	}
	close()
}
