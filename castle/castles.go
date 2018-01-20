package castle

import (
	"github.com/tamura2004/nobugo/castle/area"
)

type Castles []Castle

func Deck() Castles {
	return Castles{
		New("羅馬", "法王庁", 1, 1, 1, area.W6),
		New("天竺", "東インド会社", 2, 3, 0, area.W5),
		New("唐", "万里の長城", 6, 0, 0, area.W5),
		New("琉球", "首里城", 2, 0, 0, area.W4),
		New("薩摩", "鹿児島城", 5, 0, 0, area.W4),
		New("土佐", "高知城", 5, 0, 0, area.W3),
		New("出雲", "出雲大社", 1, 1, 0, area.W3),
		New("備前", "岡山城", 2, 0, 0, area.W2),
		New("播磨", "姫路城", 3, 4, 0, area.W2),
		New("河内", "石山本願寺", 5, 0, 0, area.W1),
		New("大和", "東大寺", 1, 1, 0, area.W1),
		New("山城", "二条城", 1, 1, 1, area.W1),
		New("美濃", "稲葉山城", 5, 0, 0, area.E1),
		New("三河", "岡崎城", 2, 0, 0, area.E1),
		New("尾張", "名古屋城", 4, 0, 0, area.E1),
		New("駿河", "駿府城", 3, 0, 0, area.E2),
		New("相模", "小田原城", 2, 3, 4, area.E2),
		New("甲斐", "躑躅ヶ崎館", 5, 0, 0, area.E2),
		New("越後", "春日山城", 2, 3, 0, area.E3),
		New("武蔵", "江戸城", 3, 4, 0, area.E3),
		New("出羽", "米沢城", 3, 0, 0, area.E4),
		New("蝦夷", "五稜郭", 5, 0, 0, area.E4),
		New("桜面戸", "金門橋", 1, 1, 0, area.E5),
		New("紐育", "萬八端島", 3, 4, 5, area.E6),
	}
}

func (cd Castles) Value() (a []string) {
	for _, c := range cd {
		a = append(a, c.Name)
	}
	return a
}
