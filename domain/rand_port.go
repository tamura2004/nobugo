package domain

var Rand interface {
	Intn(int) int
	Seed(int64)
}
