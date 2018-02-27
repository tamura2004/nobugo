package adapter

import (
	"fmt"
	"github.com/tamura2004/nobugo/domain"
)

type DicePresenter domain.Dice

func (p DicePresenter) Name() string {
	return fmt.Sprintf("[%s%d]", p.Color, p.Num)
}
