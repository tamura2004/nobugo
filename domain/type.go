package domain

import "strconv"

type Type int

const (
	EMPTY Type = iota
	BUSYO
	SAMURAI
	CASTLE
	HONOJI
)

const _Type_name = "EMPTYBUSYOSAMURAICASTLEHONOJI"

var _Type_index = [...]uint8{0, 5, 10, 17, 23, 29}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
