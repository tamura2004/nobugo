package infra_test

import (
	"github.com/tamura2004/nobugo/infra"
)

func ExampleUiMsgBox() {
	infra.MsgBox("hello")
	// Output:
	// +-------+
	// | hello |
	// +-------+
}
