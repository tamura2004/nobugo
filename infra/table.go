package infra

import (
	"github.com/olekukonko/tablewriter"
	"github.com/tamura2004/nobugo/adapter"
	"os"
)

type TablePrinter struct{}

func (TablePrinter) Print(data [][]string) {
	w := newWriter()
	w.AppendBulk(data)
	w.Render()
}

func newWriter() *tablewriter.Table {
	w := tablewriter.NewWriter(os.Stdout)
	w.SetAutoMergeCells(true)
	w.SetRowLine(true)
	return w
}

func InitTablePrinter() {
	adapter.InitTablePrinter(TablePrinter{})
}
