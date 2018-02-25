package adapter

type Printable interface {
	Header() []string
	Encode() Dic
}

func Print(p Printable) {
	dic := p.Encode()
	keys := p.Header()
	table := ConvertDicToTable(dic, keys)
	TableDriver.Print(table.Data)
}
