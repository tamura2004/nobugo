package mock

type Rand struct{}

var x int

func (Rand) Intn(n int) int {
	m := x % n
	x++
	return m
}

func (Rand) Seed(n int64) {
	x = int(n)
}
