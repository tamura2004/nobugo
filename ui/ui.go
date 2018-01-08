package ui

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/peterh/liner"
)

func Pause(msg string) {
	l := liner.NewLiner()
	defer l.Close()
	l.Prompt(msg)
}

func MsgBox(msg string) {
	t := tablewriter.NewWriter(os.Stdout)
	t.Append([]string{msg})
	t.Render()
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
