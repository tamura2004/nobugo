package pool

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Poolはダイスの数と振った結果を保有する
type Pool struct {
	Rolled bool   //振られている状態であればtrue
	Num    int    //ダイスの個数
	Dice   [6]int //添え字が出目、値が出たダイスの数
}

func New(n int) *Pool {
	return &Pool{
		Num:  n,
		Dice: [6]int{},
	}
}

// RollはPoolのダイスをすべて振る
func (p *Pool) Roll() {
	p.Dice = [6]int{}
	for i := 0; i < p.Num; i++ {
		p.Dice[d6()]++
	}
	p.Rolled = true
}

//
func (p *Pool) Selection() ([]string, [][2]int) {
	items := []string{}
	vals := [][2]int{} // dice and value
	j := 0

	for i := 0; i < 6; i++ {
		if p.Dice[i] > 0 {
			items = append(items, fmt.Sprintf("use [%d]", i+1))
			vals = append(vals, [2]int{i, p.Dice[i]})
			j++
		}
	}

	return items, vals
}

func (p *Pool) Value() []string {
	if p.Rolled {
		return []string{
			p.String(),
		}
	} else {
		return []string{
			strconv.Itoa(p.Num),
		}
	}
}

func (p *Pool) String() string {
	s := ""
	for i := 5; i >= 0; i-- {
		for j := 0; j < p.Dice[i]; j++ {
			s += fmt.Sprintf("[%d]", i+1)
		}
	}
	return s
}

// d6は0-5の出目をランダムに返す
func d6() int {
	return rand.Intn(6)
}
