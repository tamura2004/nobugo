package infra_test

import (
	"github.com/tamura2004/nobugo/infra"
)

func ExampleUiMsgBox() {
	ui := infra.UI{}
	ui.MsgBox("hello")
	// Output:
	// +-------+
	// | hello |
	// +-------+
}
