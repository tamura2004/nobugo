package adapter_test

import (
	"github.com/tamura2004/nobugo/adapter"
	"testing"
)

func NewDic() adapter.Dic {
	return []map[string][]string{
		{
			"a": {"1", "2", "3", "4"},
			"b": {"5", "6", "7"},
		},
		{
			"a": {"8"},
			"b": {"9", "0"},
		},
	}
}

func TestWidth(t *testing.T) {
	dic := NewDic()
	want := 2
	got := dic.Width()

	if want != got {
		t.Errorf("bad dic width, got %s want %s", got, want)
	}
}

func TestHeight(t *testing.T) {
	dic := NewDic()
	want := 4
	got := dic.Height("a")

	if want != got {
		t.Errorf("bad dic width, got %s want %s", got, want)
	}
}
