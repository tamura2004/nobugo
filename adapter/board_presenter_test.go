package adapter_test

import (
	"github.com/kr/pretty"
	"github.com/tamura2004/nobugo/adapter"
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/infra"
)

func init() {
	adapter.TableDriver = infra.Table{}
}

func ExampleBoardPresenterHeader() {
	b := adapter.BoardPresenter(*domain.NewBoard())
	pretty.Print(b.Header())
	// Output:
	// []string{"番号", "名称", "カード", "ダイス"}
}

func ExampleBoardPresenterPrint() {
	board := domain.NewBoard()
	dst := board.GetPool(6)
	src := domain.NewPool(domain.RED, 6)
	src.Move(dst, 0)
	b := adapter.BoardPresenter(*board)
	adapter.Print(b)
	// Output:
	// +------+------+------+------+------+------+------+
	// | 番号 |    1 |    2 |    3 |    4 |    5 |    6 |
	// +------+------+------+------+------+------+------+
	// | 名称 | 調略 | 調略 | 合戦 | 合戦 | 合戦 | 合戦
	// +------+------+------+------+------+------+------+
}
