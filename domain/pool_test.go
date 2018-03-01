package domain_test

import (
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/infra/mock"
	"testing"
)

func init() {
	domain.Rand = mock.Rand{}
}

func TestNewPool(t *testing.T) {
	p := *domain.NewPool(domain.RED, 10)
	for i := 0; i < 10; i++ {
		got := p[i].Color
		want := domain.RED

		if got != want {
			t.Errorf("bad new pool want  %q, got %q", want, got)
		}
	}
}

func TestPoolRoll(t *testing.T) {
	p := *domain.NewPool(domain.RED, 3)
	domain.Rand.Seed(0)
	p.Roll() // [1,2,3]

	for i, v := range []int{1, 2, 3} {
		if p[i].Num != v {
			t.Errorf("bad new pool roll, want %v got %v", v, p[i].Num)
		}
	}
}

func TestPoolMove(t *testing.T) {
	src := *domain.NewPool(domain.RED, 3)
	domain.Rand.Seed(0)
	src.Roll() // [1,2,3]

	dst := domain.NewPool(domain.RED, 3)
	dst.Roll() // [4,5,6]

	src.Move(dst, 2) // [1,3] [4,5,6,2]

	cases := []struct {
		i int
		v int
	}{
		{0, 4},
		{1, 5},
		{2, 6},
		{3, 2},
	}

	for _, c := range cases {
		got := (*dst)[c.i].Num
		if got != c.v {
			t.Errorf("bad pool move, want %v got %v", c.v, got)
		}
	}
}

func TestPoolReplace(t *testing.T) {
	p := *domain.NewPool(domain.RED, 3)
	domain.Rand.Seed(0)
	p.Roll()        // [1,2,3]
	p[0].Num = 2    // [2,2,3]
	p.Replace(2, 5) // [5,2,3]

	for i, v := range []int{5, 2, 3} {
		if p[i].Num != v {
			t.Errorf("bad new pool replace, want %v got %v", v, p[i].Num)
		}
	}
}

func TestPoolInclude(t *testing.T) {
	p := *domain.NewPool(domain.RED, 3)
	domain.Rand.Seed(0)
	p.Roll() // [1,2,3]

	cases := []struct {
		num     int
		include bool
	}{
		{-1, false},
		{0, false},
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, false},
		{6, false},
		{7, false},
	}

	for _, c := range cases {
		if p.Include(c.num) != c.include {
			t.Errorf("bad pool include, %v include %v is %v", p, c.num, c.include)
		}
	}
}

func TestPoolWinner(t *testing.T) {
	p := domain.Pool{}
	p.Add(domain.NewDice(domain.BLUE))
	p.Add(domain.NewDice(domain.RED))
	p.Add(domain.NewDice(domain.RED))
	p.Add(domain.NewDice(domain.BLUE))

	got := p.Winner()
	want := domain.BLUE
	if got != want {
		t.Errorf("bad pool winner got %s want %s", got, want)
	}
}
