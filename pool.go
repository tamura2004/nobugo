package main

import (
	"math/rand"
)

// Poolは6面ダイスのダイスプールを表す
type Pool []int

// Roll(n)はダイスをn個振り、大きい順に先頭から並べる
func (p *Pool) Roll(n int) {
	x := make([]int, n)
	for i, _ := range x {
		x[i] = d6()
		for j := range x[:i] {
			if x[j] < x[i] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	*p = x
}

// d6は1-6の出目をランダムに返す
func d6() int {
	return rand.Intn(6) + 1
}
