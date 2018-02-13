package infra

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type Table struct{}

func (Table) Print(data [][]string) {
	w := NewWriter()
	w.AppendBulk(data)
	w.Render()
}

func NewWriter() *tablewriter.Table {
	w := tablewriter.NewWriter(os.Stdout)
	w.SetAutoMergeCells(true)
	w.SetRowLine(true)
	return w
}
