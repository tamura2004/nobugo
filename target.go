package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Action struct {
	Owner Owner
	Bid   [MAX_PLAYER]Bid
}

const (
	MAX_PLAYER = 6
	MAX_ACTION = 6
)

type Bid struct {
	Num      int
	Priority int
}

var action struct {
	Priority int
	Deck     [MAX_ACTION]*Action
}

func init() {
	for i := 0; i < MAX_ACTION; i++ {
		action.Deck[i] = new(Action)
	}
	action.Deck[0].Bid[0].Num = 3
}

func SetupActionSamurai() {
	deck := SelectSamurai(func(s *Samurai) bool {
		return s.Owner == DECK
	})
	deck[0].Owner = Owner(ACTION1)
	deck[1].Owner = Owner(ACTION2)
}

func SetupActionCastle() {
	deck := SelectCastle(func(c *Castle) bool {
		return c.Owner == DECK
	})
	deck[0].Owner = Owner(ACTION3)
	deck[1].Owner = Owner(ACTION4)
	deck[2].Owner = Owner(ACTION5)
	deck[3].Owner = Owner(ACTION6)
}

func MaxBidRow() int {
	maxBidRow := 0
	for _, a := range action.Deck {
		bidRow := 0
		for _, b := range a.Bid {
			if b.Num != 0 {
				bidRow++
			}
		}
		if maxBidRow < bidRow {
			maxBidRow = bidRow
		}
	}
	return maxBidRow
}

func PrintAction() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetRowLine(true)
	t.SetAutoMergeCells(true)

	data := make([][]string, 5+MaxBidRow())

	for i := 0; i < len(data); i++ {
		data[i] = make([]string, 7)
	}

	data[0][0] = "アクション"
	data[1][0] = "アクション"
	data[2][0] = "カード"
	data[3][0] = "カード"
	data[4][0] = "カード"
	for i := 0; i < MaxBidRow(); i++ {
		data[i+5][0] = "ダイス"
	}

	aligns := make([]int, 7)
	aligns[0] = tablewriter.ALIGN_DEFAULT

	for i, a := range action.Deck {
		t.SetColMinWidth(i+1, 10)
		aligns[i+1] = tablewriter.ALIGN_CENTER
		if i < 2 {
			data[0][i+1] = "調略"
		} else {
			data[0][i+1] = "合戦"
		}
		data[1][i+1] = fmt.Sprint(i)

		if i < 2 {
			if s := a.Samurai(); s != nil {
				data[2][i+1] = s.Name
				data[3][i+1] = s.Ability.String()
			}
		} else {
			if c := a.Castle(); c != nil {
				data[2][i+1] = c.Name
				data[3][i+1] = c.Ability.String()
				data[4][i+1] = fmt.Sprintf("%s/%s", c.Area.String(), c.Country)
			}
		}

	}

	t.SetColumnAlignment(aligns)
	t.AppendBulk(data)
	t.Render()
}

func (a *Action) Samurai() *Samurai {
	deck := SelectSamurai(func(s *Samurai) bool {
		return s.Owner == a.Owner
	})
	if len(deck) == 0 {
		return nil
	}
	return deck[0]
}

func (a *Action) Castle() *Castle {
	deck := SelectCastle(func(c *Castle) bool {
		return c.Owner == a.Owner
	})
	if len(deck) == 0 {
		return nil
	}
	return deck[0]
}
