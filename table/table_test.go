package table

type tableTest struct{}

func (t *tableTest) Values() []map[string][]string {
	return []map[string][]string{
		map[string][]string{
			"アクション": []string{"調略", "1"},
			"カード":   []string{"伊藤マンショ", "西６"},
			"ダイス":   []string{"黒２", "赤３"},
		},
		map[string][]string{
			"アクション": []string{"調略", "2"},
			"カード":   []string{"木下藤吉郎", "[1,2]->6"},
		},
		map[string][]string{
			"アクション": []string{"合戦", "3"},
			"カード":   []string{"東インド会社", "[5][4][3]", "E1/桜面戸"},
		},
		map[string][]string{
			"アクション": []string{"合戦", "4"},
			"カード":   []string{"万里の長城", "[4][3]", "W5/唐"},
			"ダイス":   []string{"白2", "黒1", "赤2"},
		},
		map[string][]string{
			"アクション": []string{"合戦", "5"},
			"カード":   []string{"姫路城", "[5]", "W1/播磨"},
		},
		map[string][]string{
			"アクション": []string{"合戦", "6"},
			"カード":   []string{"稲葉山城", "[4]", "E1/美濃"},
		},
	}
}

func (t *tableTest) render() {
	Render(t, []string{"アクション", "カード", "ダイス"})
}

func ExampleRender() {
	t := &tableTest{}
	t.render()
	// Output:
	//+------------+--------------+------------+--------------+------------+---------+----------+
	//| アクション | 調略         | 調略       | 合戦         | 合戦       | 合戦    | 合戦     |
	//+            +--------------+------------+--------------+------------+---------+----------+
	//|            |            1 |          2 |            3 |          4 |       5 |        6 |
	//+------------+--------------+------------+--------------+------------+---------+----------+
	//| カード     | 伊藤マンショ | 木下藤吉郎 | 東インド会社 | 万里の長城 | 姫路城  | 稲葉山城 |
	//+            +--------------+------------+--------------+------------+---------+----------+
	//|            | 西６         | [1,2]->6   | [5][4][3]    | [4][3]     | [5]     | [4]      |
	//+            +--------------+------------+--------------+------------+---------+----------+
	//|            |              |            | E1/桜面戸    | W5/唐      | W1/播磨 | E1/美濃  |
	//+------------+--------------+------------+--------------+------------+---------+----------+
	//| ダイス     | 黒２         |            |              | 白2        |         |          |
	//+            +--------------+------------+--------------+------------+---------+----------+
	//|            | 赤３         |            |              | 黒1        |         |          |
	//+            +--------------+------------+--------------+------------+---------+----------+
	//|            |              |            |              | 赤2        |         |          |
	//+------------+--------------+------------+--------------+------------+---------+----------+
}
