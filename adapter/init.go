package adapter

import (
	"github.com/tamura2004/nobugo/usecase"
	"github.com/tamura2004/nobugo/usecase/port"
)

func Init() {
	port.Party = &Party{}
	port.Board = &Party{}

	usecase.Init()
}
