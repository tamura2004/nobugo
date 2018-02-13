package usecase

import (
	"github.com/tamura2004/nobugo/usecase/port"
)

var step Step

func Run() {
	open()
	for next() {
		port.Party.Print()
		port.Board.Print()
	}
	close()
}

func next() bool {
	switch step {

	case PREPARE:
		return prepare()

	case MARCH:
		march()

	case EMPLOY:
		employ()

	case BATTLE:
		battle()

	case CHECK:
		check()
	}
	return false
}
