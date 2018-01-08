package main

import "testing"

// 結果が1-6の整数であること
func TestD6(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := d6(); n < 1 || 6 < n {
			t.Errorf("out of range d6 = %d", n)
		}
	}
}

// 結果が降順にソートされていること
func TestPoolRoll(t *testing.T) {
	p := Pool{}
	n := 100
	p.Roll(n)
	for i := 0; i < int(n)-1; i++ {
		if p[i] < p[i+1] {
			t.Errorf("order reverse %d > %d", p[i], p[i+1])
		}
	}
}
