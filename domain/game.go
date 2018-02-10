package domain

type Game struct {
	Turn int
	Step
	Players
	Board
}

/*
# ゲームの準備

旗本デッキ
武将デッキ
城デッキを準備してシャッフルする
城デッキの底に、本能寺カードを入れる

プレイヤー人数を決める
各プレイヤーに旗本武将カードを１枚ずつ配る
旗本デッキを武将デッキに加えシャッフルする

投票カードを並べる

エリアカードを並べる
尾張がある東１に、信長マーカーを置く
*/
func (g Game) Open() {

}

func (g Game) Run() {

}

func (g Game) Close() {

}
