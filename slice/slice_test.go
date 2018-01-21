package slice_test

import (
	"github.com/tamura2004/nobugo/slice"
	"testing"
)

func ReverseTest(t *testing.T) {
	got := []int{1, 2, 3, 4, 5, 6, 7}
	slice.Reverse(got)

	want := []int{7, 6, 5, 4, 3, 2, 1}

	if got != want {
		t.Errorf("bad reverse got %v want %v", got, want)
	}

}
