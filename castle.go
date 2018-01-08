package main

type Castle struct {
	Owner   Owner
	Area    Area
	Country string
	Name    string
	Ability Ability
}

var castle struct {
	Deck []*Castle
}

func init() {
	castle.Deck = []*Castle{
		NewCastle("羅馬", "法王庁", 1, 1, 1, W6),
		NewCastle("天竺", "東インド会社", 2, 3, 0, W5),
		NewCastle("唐", "万里の長城", 6, 0, 0, W5),
		NewCastle("琉球", "首里城", 2, 0, 0, W4),
		NewCastle("薩摩", "鹿児島城", 5, 0, 0, W4),
		NewCastle("土佐", "高知城", 5, 0, 0, W3),
		NewCastle("出雲", "出雲大社", 1, 1, 0, W3),
		NewCastle("備前", "岡山城", 2, 0, 0, W2),
		NewCastle("播磨", "姫路城", 3, 4, 0, W2),
		NewCastle("河内", "石山本願寺", 5, 0, 0, W1),
		NewCastle("大和", "東大寺", 1, 1, 0, W1),
		NewCastle("山城", "二条城", 1, 1, 1, W1),
		NewCastle("美濃", "稲葉山城", 5, 0, 0, E1),
		NewCastle("三河", "岡崎城", 2, 0, 0, E1),
		NewCastle("尾張", "名古屋城", 4, 0, 0, E1),
		NewCastle("駿河", "駿府城", 3, 0, 0, E2),
		NewCastle("相模", "小田原城", 2, 3, 4, E2),
		NewCastle("甲斐", "躑躅ヶ崎館", 5, 0, 0, E2),
		NewCastle("越後", "春日山城", 2, 3, 0, E3),
		NewCastle("武蔵", "江戸城", 3, 4, 0, E3),
		NewCastle("出羽", "米沢城", 3, 0, 0, E4),
		NewCastle("蝦夷", "五稜郭", 5, 0, 0, E4),
		NewCastle("桜面戸", "金門橋", 1, 1, 0, E5),
		NewCastle("紐育", "萬八端島", 3, 4, 5, E6),
	}
}

func NewCastle(country, name string, x, y, z int, a Area) *Castle {
	return &Castle{
		Name:    name,
		Country: country,
		Area:    a,
		Ability: NewAbility(DEFENCE, x, y, z),
	}
}

func (c *Castle) Map() map[string]string {
	return map[string]string{
		"name":    c.Name,
		"country": c.Country,
		"area":    c.Area.String(),
		"ability": c.Ability.String(),
	}
}
