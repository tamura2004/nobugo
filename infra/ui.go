package infra

import (
	"fmt"
	"github.com/peterh/liner"
	"github.com/tamura2004/nobugo/domain"
	"os"
)

type UI struct{}

func (UI) Pause(msg string) {
	l := liner.NewLiner()
	defer l.Close()
	_, err := l.Prompt(msg)
	if err != nil {
		os.Exit(0)
	}
}

func (UI) MsgBox(msg string) {
	t := Table{}
	t.Print([][]string{{msg}})
}

func (UI) Num(min, max int, msg string) int {
	var n int
	for {
		fmt.Printf("%s(%d-%d)? >", msg, min, max)
		fmt.Scan(&n)
		if min <= n && n <= max {
			return n
		}
	}
}

func (u UI) Select(actions []domain.Action) {
	for i, action := range actions {
		fmt.Printf("%3d: %s\n", i+1, action.Msg)
	}
	n := u.Num(1, len(actions), "")
	actions[n-1].Do()
}
