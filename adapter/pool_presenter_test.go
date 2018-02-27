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

func ExamplePoolPresenterName() {
	pool := domain.NewPool(domain.RED, 3)
	domain.Rand.Seed(0)
	pool.Roll()
	dp := adapter.PoolPresenter(pool)
	pretty.Print(dp.Names())
	// Output:
	// []string{"[赤1]", "[赤2]", "[赤3]"}
}
