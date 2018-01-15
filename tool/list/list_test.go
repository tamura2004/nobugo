package list_test

import (
	"fmt"
	"github.com/tamura2004/nobugo/tool/list"
	"testing"
)

func ExamplePutNormal() {
	var ol list.List
	ol = list.New()

	ol.Put(1)
	ol.Put(2)
	ol.Put(3)
	ol.Put(2)
	ol.Put(1)

	fmt.Printf("%v", ol)
	// Output:
	// [3 2 1]
}

func ExamplePutReverse() {
	var ol list.List
	ol = list.New()

	ol.Put(1)
	ol.Put(2)
	ol.Put(3)
	ol.Put(2)
	ol.Put(1)

	fmt.Printf("%v", ol.Reverse())
	// Output:
	// [1 2 3]
}

func TestShuffle(t *testing.T) {
	ol := list.New()
	for i := 0; i < 100; i++ {
		ol.Put(i)
	}

	before := ol
	ol.Shuffle()
	after := ol

	if before[0] == after[0] {
		t.Errorf("bad shuffle, before %v after %v", before, after)
	}

}
