package usecase

import (
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/domain/entity"
)

func Init() {
	entity.Board = domain.NewBoard()
}
