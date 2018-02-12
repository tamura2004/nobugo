package adapter

type tablePrinter interface {
	Print(data [][]string)
}

var TablePrinter tablePrinter

func InitTablePrinter(t tablePrinter) {
	TablePrinter = t
}
