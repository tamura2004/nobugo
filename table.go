package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Tabled interface {
	Header() []string
	Data() [][]string
}

func Table(x Tabled) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader(x.Header())
	t.AppendBulk(x.Data())
	t.Render()
}
