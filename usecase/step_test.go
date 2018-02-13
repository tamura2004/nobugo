package usecase_test

import (
	"github.com/tamura2004/nobugo/usecase"
	"testing"
)

var got usecase.Step

func TestStepNext(t *testing.T) {
	want := usecase.PREPARE

	got = usecase.CHECK
	got.Next()

	if want != got {
		t.Errorf("bad step next(), got %s want %s", got, want)
	}

}
