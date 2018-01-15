package pool

import (
	"fmt"
	"testing"
)

func TestRoll(t *testing.T) {
	var pool *Pool
	const NUM = 10
	pool = New(NUM)

	if pool.Num != NUM {
		t.Errorf("got %v want %d", pool.Num, NUM)
	}

	pool.Roll()

	if !pool.Rolled {
		t.Errorf("got Rolled false want true")
	}

	sum := 0
	for i := 0; i < 6; i++ {
		sum += pool.Dice[i]
	}

	if sum != NUM {
		t.Errorf("got dice sum = %v want %v", sum, NUM)
	}
}

func ExampleString() {
	pool := New(6)
	pool.Dice[5] = 3
	pool.Dice[4] = 2
	pool.Dice[3] = 1
	fmt.Println(pool.String())
	// Output:
	// [6][6][6][5][5][4]
}
