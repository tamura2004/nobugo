package domain_test

import (
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/infra/mock"
	"testing"
)

func init() {
	domain.Rand = mock.Rand{}
}

func TestNewDice(t *testing.T) {
	d := domain.NewDice(domain.RED)

	if d.Num != 0 {
		t.Errorf("bad new dice num %q", d.Num)
	}

	if d.Color != domain.RED {
		t.Errorf("bad new dice color %q", d.Color)
	}
}

func TestDiceRoll(t *testing.T) {
	d := domain.NewDice(domain.RED)
	d.Roll()

	if d.Num != 1 {
		t.Errorf("bad new dice roll num %q", d.Num)
	}

	if d.Color != domain.RED {
		t.Errorf("bad new dice color %q", d.Color)
	}
}

func TestDiceRestore(t *testing.T) {
	d := domain.NewDice(domain.RED)
	d.Roll()
	d.Restore()

	if d.Num != 0 {
		t.Errorf("bad new dice num %q", d.Num)
	}

	if d.Color != domain.RED {
		t.Errorf("bad new dice color %q", d.Color)
	}
}
