package infra

import (
	"math/rand"
)

type Rand struct{}

func (Rand) Intn(n int) int {
	return rand.Intn(n)
}

func (Rand) Seed(n int64) {
	rand.Seed(n)
}
