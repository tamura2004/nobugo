package infra

import (
	"fmt"

	"github.com/peterh/liner"
)

func Pause(msg string) {
	l := liner.NewLiner()
	defer l.Close()
	l.Prompt(msg)
}

func MsgBox(msg string) {
	t := TablePrinter{}
	t.Print([][]string{{msg}})
}

func InputNumber(msg string, min, max int) (n int) {
	for {
		fmt.Printf("%s(%d-%d)? >", msg, min, max)
		fmt.Scan(&n)
		if min <= n && n <= max {
			return n
		}
	}
}

func SelectNumber(msg string, items []string) (n int) {
	for i, item := range items {
		fmt.Printf("%3d: %s\n", i+1, item)
	}
	return InputNumber(msg, 1, len(items))
}
