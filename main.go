package main

import (
	"github.com/tamura2004/nobugo/game"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	g := game.New()
	for g.Next() {
	}
}
