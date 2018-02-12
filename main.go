package main

import (
	"fmt"
	"github.com/tamura2004/nobugo/domain"
	"github.com/tamura2004/nobugo/usecase"
)

type printer struct{}

func (p printer) Print() {
	fmt.Println("running!")
}

func main() {
	g := usecase.Game{
		domain.Game{},
		printer{},
	}

	g.Run()
}
