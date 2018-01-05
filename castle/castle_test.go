package castle_test

import (
	"github.com/tamura2004/nobugo/castle"
)

func ExamplePrintTable() {
	x := castle.Deck()
	x.PrintTable()
	// Output:
	//+------+------+--------+--------------+---------+
	//| 番号 | 地域 |   国   |      城      |  防御   |
	//+------+------+--------+--------------+---------+
	//|    1 | 西6  | 羅馬   | 法王庁       | [1 1 1] |
	//|    2 | 西5  | 天竺   | 東インド会社 | [2 3]   |
	//|    3 | 西5  | 唐     | 万里の長城   | [6]     |
	//|    4 | 西4  | 琉球   | 首里城       | [2]     |
	//|    5 | 西4  | 薩摩   | 鹿児島城     | [5]     |
	//|    6 | 西3  | 土佐   | 高知城       | [5]     |
	//|    7 | 西3  | 出雲   | 出雲大社     | [1 1]   |
	//|    8 | 西2  | 備前   | 岡山城       | [2]     |
	//|    9 | 西2  | 播磨   | 姫路城       | [3 4]   |
	//|   10 | 西1  | 河内   | 石山本願寺   | [5]     |
	//|   11 | 西1  | 大和   | 東大寺       | [1 1]   |
	//|   12 | 西1  | 山城   | 二条城       | [1 1 1] |
	//|   13 | 東1  | 美濃   | 稲葉山城     | [5]     |
	//|   14 | 東1  | 三河   | 岡崎城       | [2]     |
	//|   15 | 東1  | 尾張   | 名古屋城     | [4]     |
	//|   16 | 東2  | 駿河   | 駿府城       | [3]     |
	//|   17 | 東2  | 相模   | 小田原城     | [2 3 4] |
	//|   18 | 東2  | 甲斐   | 躑躅ヶ崎館   | [5]     |
	//|   19 | 東3  | 越後   | 春日山城     | [2 3]   |
	//|   20 | 東3  | 武蔵   | 江戸城       | [3 4]   |
	//|   21 | 東4  | 出羽   | 米沢城       | [3]     |
	//|   22 | 東4  | 蝦夷   | 五稜郭       | [5]     |
	//|   23 | 東5  | 桜面戸 | 金門橋       | [1 1]   |
	//|   24 | 東6  | 紐育   | 萬八端島     | [3 4 5] |
	//+------+------+--------+--------------+---------+
}
