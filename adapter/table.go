package adapter

import (
	"log"
)

type Table struct {
	Width  int
	Height int
	Data   [][]string
}

func NewTable(w, h int) *Table {
	t := &Table{
		Width:  w,
		Height: h,
	}

	for i := 0; i < h; i++ {
		t.Data = append(t.Data, make([]string, w))
	}

	return t
}

func NewRows(dic Dic, key string) *Table {
	t := NewTable(dic.Width()+1, dic.Height(key))

	for y := 0; y < t.Height; y++ {
		t.Data[y][0] = key
	}

	for x, col := range dic {
		for y, r := range col[key] {
			t.Data[y][x+1] = r
		}
	}

	return t
}

func (t *Table) Append(o *Table) {
	if t.Width != o.Width {
		log.Fatal("width mismatch %d != %d", t.Width, o.Width)
	}
	t.Height += o.Height
	for _, row := range o.Data {
		t.Data = append(t.Data, row)
	}
}

func ConvertDicToTable(dic Dic, keys []string) *Table {
	t := NewTable(dic.Width()+1, 0)

	for _, key := range keys {
		rows := NewRows(dic, key)
		t.Append(rows)
	}

	return t
}
