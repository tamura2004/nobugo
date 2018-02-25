package usecase

import (
	"github.com/tamura2004/nobugo/domain"
)

var UI interface {
	Num(min, max int, msg string) int
	Select([]domain.Action)
	Pause(msg string)
	MsgBox(msg string)
}
