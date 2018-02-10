package usecase

import (
	"github.com/github.com/tamura2004/nobugo/domain"
)

func Run() {
	initGame()
	for next() {}
	closeGame()
}



var game domain.Game

type Printer interface {
	Print()
}

var pr struct {
	title Printer
}

type Inputer interface {
	InputNum(msg string) int
	InputRangeNum(msg string min,max int) int
}

var in struct {

}

type RangeScanner interface {
	Scan(int, int, string) int
}

type PlayerGroup interface {
	Add()
	AddN(int)
}

func Run() {
	checkInitErr()
	game.Open()

}

type Deck interface {
	Draw() Card
	DrawN(n int) Deck
}

var deck struct {
	Samurai Deck
	Castle  Deck
}

func (g *domain.Game) Open() {
	CheckInitErr()

	TitlePrinter.Print()
	n := RangeScanner.Scan(3, 6, "プレイヤー人数？")
	PlayerGroup.AddN(n)
	PlayerGroup.Each(func(p *domain.Player) {
		s := SamuraiDeck.Draw()
		p.AddSamurai(s)
	})
}

func checkInitErr() {

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

# 準備フェイズ
投票カード１，２にそれぞれ武将カードを配置する
城カードを6枚引く、本能寺カードを引いたら、最終ターンを行う

一番東と西のカードを、デッキの底に戻す
残りの城カードを、投票カード３～６に並べる

各プレイヤーは、ダイス３つに加え、武将と城のペア事にダイスを１つえる。
ダイスの上限は１０個

# 行軍フェイズ
開始プレイヤーはダイスを１つ追加で得る

## 各プレイヤーの処理
手番プレイヤーはダイスをすべて振る
プレイヤーは武将能力を使用してもよい
プレイヤーはダイス目を一つ選び、該当するすべてのダイスを投票カードに置く
または一向宗が支配する城の防御力に該当する目
次のプレイヤーを手番プレイヤーにする

全てのプレイヤーの手元ダイスがなくなるまで繰り返す

# 調略フェイズ
投票カード１、２の武将カードについて、最も多くのダイスを置いたプレイヤーが得る。
ダイスの数が同じであれば、後からおいたプレイヤーが優先

# 合戦フェイズ
投票カード３～６の城カードについて、最も多くのダイスを置いたプレイヤーが総大将となる。
ダイスの数が同じであれば後から置いたプレイヤーが優先。
総大将は、他プレイヤーのダイスを含めすべてを振り直す。

城カードのエリアが、信長マーカーまたは城マーカーと隣接していないなら、間のエリア一つ
毎に、ダイスを１つ取り除く（補給線）

残ったダイスについて、城カードの防御力以上の目を割り当てられれば城カードを得る。
割り当てられないなら、一向宗が城を得る

# 終了フェイズ
城に武将を割り当てる









*/
