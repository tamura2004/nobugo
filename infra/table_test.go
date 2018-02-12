package infra_test

import (
	// "github.com/tamura2004/nobugo/adapter"
	"github.com/tamura2004/nobugo/infra"
)

func ExampleTablePrint() {
	var p infra.TablePrinter

	data := [][]string{
		{"a", "b", "c"},
		{"a", "d", "e"},
		{"f", "b", "c"},
		{"f", "d", "e"},
	}

	p.Print(data)
	// Output:
	// +---+---+---+
	// | a | b | c |
	// +   +---+---+
	// |   | d | e |
	// +---+---+---+
	// | f | b | c |
	// +   +---+---+
	// |   | d | e |
	// +---+---+---+
}
