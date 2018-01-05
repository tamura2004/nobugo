package castle

import (
	"fmt"

	"github.com/tamura2004/nobugo/deck"
)

type Area struct {
	isEast   bool
	distance int
}

func (a Area) String() string {
	if a.isEast {
		return fmt.Sprintf("東%d", a.distance)
	} else {
		return fmt.Sprintf("西%d", a.distance)
	}
}

type Castle struct {
	number  int
	country string
	name    string
	defence []int
	Area
}

func (c Castle) String() string {
	return c.name
}

func (c Castle) Header() []string {
	return []string{"番号", "地域", "国", "城", "防御"}
}

func (c Castle) Values() []string {
	return []string{
		fmt.Sprint(c.number),
		c.Area.String(),
		c.country,
		c.name,
		fmt.Sprintf("%v", c.defence),
	}
}

func Deck() deck.Deck {
	return deck.Deck{
		Castle{1, "羅馬", "法王庁", []int{1, 1, 1}, Area{false, 6}},
		Castle{2, "天竺", "東インド会社", []int{2, 3}, Area{false, 5}},
		Castle{3, "唐", "万里の長城", []int{6}, Area{false, 5}},
		Castle{4, "琉球", "首里城", []int{2}, Area{false, 4}},
		Castle{5, "薩摩", "鹿児島城", []int{5}, Area{false, 4}},
		Castle{6, "土佐", "高知城", []int{5}, Area{false, 3}},
		Castle{7, "出雲", "出雲大社", []int{1, 1}, Area{false, 3}},
		Castle{8, "備前", "岡山城", []int{2}, Area{false, 2}},
		Castle{9, "播磨", "姫路城", []int{3, 4}, Area{false, 2}},
		Castle{10, "河内", "石山本願寺", []int{5}, Area{false, 1}},
		Castle{11, "大和", "東大寺", []int{1, 1}, Area{false, 1}},
		Castle{12, "山城", "二条城", []int{1, 1, 1}, Area{false, 1}},
		Castle{13, "美濃", "稲葉山城", []int{5}, Area{true, 1}},
		Castle{14, "三河", "岡崎城", []int{2}, Area{true, 1}},
		Castle{15, "尾張", "名古屋城", []int{4}, Area{true, 1}},
		Castle{16, "駿河", "駿府城", []int{3}, Area{true, 2}},
		Castle{17, "相模", "小田原城", []int{2, 3, 4}, Area{true, 2}},
		Castle{18, "甲斐", "躑躅ヶ崎館", []int{5}, Area{true, 2}},
		Castle{19, "越後", "春日山城", []int{2, 3}, Area{true, 3}},
		Castle{20, "武蔵", "江戸城", []int{3, 4}, Area{true, 3}},
		Castle{21, "出羽", "米沢城", []int{3}, Area{true, 4}},
		Castle{22, "蝦夷", "五稜郭", []int{5}, Area{true, 4}},
		Castle{23, "桜面戸", "金門橋", []int{1, 1}, Area{true, 5}},
		Castle{24, "紐育", "萬八端島", []int{3, 4, 5}, Area{true, 6}},
	}
}
