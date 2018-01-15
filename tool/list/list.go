package list

import (
	"fmt"
	"math/rand"
)

type List []interface{}

func New() List {
	return List{}
}

func (ol *List) Put(el interface{}) {
	if ix := ol.find(el); ix != -1 {
		ol.Delete(ix)
	}
	*ol = append(*ol, el)
}

func (ol List) Reverse() List {
	list := make(List, len(ol))
	for i, e := range ol {
		list[len(ol)-i-1] = e
	}
	return list
}

func (ol *List) find(el interface{}) int {
	for i, e := range *ol {
		if el == e {
			return i
		}
	}
	return -1
}

func (ol *List) Delete(ix int) interface{} {
	if ix < 0 || len(*ol) <= ix {
		fmt.Printf("out of range, ix = %v\n", ix)
		return nil
	}

	e := (*ol)[ix]
	*ol = append((*ol)[:ix], (*ol)[ix+1:]...)
	return e
}

func (ol *List) Shuffle() {
	x := make(List, len(*ol))
	for i, ix := range rand.Perm(len(*ol)) {
		x[i] = (*ol)[ix]
	}
	*ol = x
}
