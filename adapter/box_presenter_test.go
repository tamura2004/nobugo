package adapter_test

import (
	"github.com/kr/pretty"
	"github.com/tamura2004/nobugo/adapter"
	"github.com/tamura2004/nobugo/domain"
)

func ExampleBoxPresenterMap() {
	box := domain.NewBox(2)
	pretty.Print(adapter.BoxPresenter(box).Map()["名称"])
	// Output:
	// []string{"合戦"}
}
