package list

import (
	"fmt"
	"math/rand"
)

type List []Element

type Element interface {
	Title() string
	Detail() []string
}

func New() List {
	return List{}
}

func (ol *List) Put(el Element) {
	if ix := ol.find(el); ix != -1 {
		ol.Delete(ix)
	}
	*ol = append(*ol, el)
}

func (ol *List) find(el Element) int {
	for i, e := range *ol {
		if el == e {
			return i
		}
	}
	return -1
}

func (ol *List) Delete(ix int) Element {
	if ix < 0 || len(*ol) <= ix {
		fmt.Printf("out of range, ix = %v\n", ix)
		return nil
	}

	e := (*ol)[ix]
	*ol = append((*ol)[:ix], (*ol)[ix+1:]...)
	return e
}

func (ol *List) Reverse() {
	a := *ol
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func (ol *List) Shuffle() {
	x := make(List, len(*ol))
	for i, ix := range rand.Perm(len(*ol)) {
		x[i] = (*ol)[ix]
	}
	*ol = x
}
