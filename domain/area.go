package domain

import (
	"fmt"
	"strings"
)

type Area int

const (
	W6 Area = iota + 1
	W5
	W4
	W3
	W2
	W1
	E1
	E2
	E3
	E4
	E5
	E6
)

const AREA_STRING = "西6西5西4西3西2西1東1東2東3東4東5東6"

var AREA_INDEX = []int{0, 4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48}

func (a Area) String() string {
	ix := AREA_INDEX

	if int(a) < 1 || len(ix) < int(a) {
		return "不明なエリア"
	}

	return AREA_STRING[ix[a-1]:ix[a]]
}

func (a *Area) UnmarshalText(text []byte) error {
	offset := strings.Index(AREA_STRING, string(text))

	for i, ix := range AREA_INDEX {
		if offset == ix {
			*a = Area(i + 1)
			return nil
		}
	}
	return fmt.Errorf("bad area string %s", string(text))
}
