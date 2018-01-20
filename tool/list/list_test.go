package list_test

import (
	"fmt"
	"github.com/tamura2004/nobugo/tool/list"
	"strconv"
	"testing"
)

type element int

func (e element) String() string {
	return strconv.Itoa(int(e))
}

func (e element) Value() []string {
	return []string{
		e.String(),
	}
}

func ExamplePutNormal() {
	var ol list.List
	ol = list.New()

	ol.Put(element(1))
	ol.Put(element(2))
	ol.Put(element(3))
	ol.Put(element(2))
	ol.Put(element(1))

	fmt.Printf("%v", ol)
	// Output:
	// [3 2 1]
}

func ExamplePutReverse() {
	var ol list.List
	ol = list.New()

	ol.Put(element(1))
	ol.Put(element(2))
	ol.Put(element(3))
	ol.Put(element(2))
	ol.Put(element(1))

	ol.Reverse()
	fmt.Printf("%v", ol)
	// Output:
	// [1 2 3]
}

func TestShuffle(t *testing.T) {
	ol := list.New()
	for i := 0; i < 100; i++ {
		ol.Put(element(i))
	}

	before := ol
	ol.Shuffle()
	after := ol

	if before[0] == after[0] {
		t.Errorf("bad shuffle, before %v after %v", before, after)
	}

}
