package main

import "strconv"

const _Color_name = "黒白赤青緑黄"

var _Color_index = [...]uint8{0, 0, 3, 6, 9, 12, 15, 18}

func (i Color) String() string {
	if i < 0 || i >= Color(len(_Color_index)-1) {
		return "Color(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Color_name[_Color_index[i]:_Color_index[i+1]]
}
