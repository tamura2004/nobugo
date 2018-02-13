package main

import (
	"github.com/tamura2004/nobugo/infra"
	"github.com/tamura2004/nobugo/usecase"
)

func main() {
	infra.Init()
	usecase.Run()
}
