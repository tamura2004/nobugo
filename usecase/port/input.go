package port

import (
	"github.com/tamura2004/nobugo/domain"
)

type input interface {
	Num(min, max int, msg string) int
	Select([]domain.Action)
	Pause(msg string)
	MsgBox(msg string)
}

var Input input
