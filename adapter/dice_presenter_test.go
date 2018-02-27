package adapter_test

import (
	"github.com/kr/pretty"
	"github.com/tamura2004/nobugo/adapter"
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/infra/mock"
)

func init() {
	domain.Rand = mock.Rand{}
}

func ExampleDicePresenterName() {
	dice := domain.NewDice(domain.RED)
	domain.Rand.Seed(0)
	dice.Roll()

	dp := adapter.DicePresenter(dice)
	pretty.Print(dp.Name())
	// Output:
	// [赤1]
}
