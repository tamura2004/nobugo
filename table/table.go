package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Tabulated interface {
	Values() []map[string][]string
}

func Render(t Tabulated, keys []string) {
	w := newWriter()
	data := [][]string{}
	for _, k := range keys {
		setRows(t, &data, k)
	}
	w.AppendBulk(data)
	w.Render()
}

func newWriter() *tablewriter.Table {
	w := tablewriter.NewWriter(os.Stdout)
	w.SetAutoMergeCells(true)
	w.SetRowLine(true)
	return w
}

func setRows(t Tabulated, data *[][]string, k string) {
	for ypos := 0; ; ypos++ {
		if row, empty := makeRow(t, ypos, k); empty {
			break
		} else {
			*data = append(*data, row)
		}
	}
}

func makeRow(t Tabulated, ypos int, k string) ([]string, bool) {
	row := make([]string, len(t.Values())+1)
	row[0] = k
	empty := true
	for col, vmap := range t.Values() {
		if ypos < len(vmap[k]) {
			row[col+1] = vmap[k][ypos]
			empty = false
		}
	}
	return row, empty
}
