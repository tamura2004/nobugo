package driver

type table interface {
	Print(data [][]string)
}

var Table table
