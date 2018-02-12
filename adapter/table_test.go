package adapter_test

import (
	"github.com/kr/pretty"
	"github.com/tamura2004/nobugo/adapter"
)

func ExampleNewTable() {
	got := adapter.NewTable(2, 3)
	pretty.Print(got)
	// Output:
	// &adapter.Table{
	//     Width:  2,
	//     Height: 3,
	//     Data:   {
	//         {"", ""},
	//         {"", ""},
	//         {"", ""},
	//     },
	// }
}

func ExampleNewRows() {
	got := adapter.NewRows(NewDic(), "b")
	pretty.Print(got)
	// Output:
	// &adapter.Table{
	//     Width:  3,
	//     Height: 3,
	//     Data:   {
	//         {"b", "5", "9"},
	//         {"b", "6", "0"},
	//         {"b", "7", ""},
	//     },
	// }
}

func ExampleAppend() {
	a := adapter.NewRows(NewDic(), "a")
	b := adapter.NewRows(NewDic(), "b")
	a.Append(b)
	pretty.Print(a)
	// Output:
	// &adapter.Table{
	//     Width:  3,
	//     Height: 7,
	//     Data:   {
	//         {"a", "1", "8"},
	//         {"a", "2", ""},
	//         {"a", "3", ""},
	//         {"a", "4", ""},
	//         {"b", "5", "9"},
	//         {"b", "6", "0"},
	//         {"b", "7", ""},
	//     },
	// }
}

func ExampleConvertDicToTable() {
	dic := NewDic()
	t := adapter.ConvertDicToTable(dic, []string{"a", "b"})
	pretty.Print(t)
	// Output:
	// &adapter.Table{
	//     Width:  3,
	//     Height: 7,
	//     Data:   {
	//         {"a", "1", "8"},
	//         {"a", "2", ""},
	//         {"a", "3", ""},
	//         {"a", "4", ""},
	//         {"b", "5", "9"},
	//         {"b", "6", "0"},
	//         {"b", "7", ""},
	//     },
	// }
}
