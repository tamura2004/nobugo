package domain_test

import (
	"github.com/tamura2004/nobugo/domain"
	"testing"
)

func TestStepNext(t *testing.T) {
	want := domain.PREPARE
	got := domain.CHECK.Next()

	if want != got {
		t.Errorf("bad step next(), got %s want %s", got, want)
	}

}
