package adapter

import (
	"github.com/tamura2004/nobugo/domain"
)

type BoardPresenter domain.Board

func (p BoardPresenter) Header() []string {
	return []string{"番号", "名称", "カード", "ダイス"}
}

func (p BoardPresenter) Encode() Dic {
	dic := Dic{}

	for _, box := range p {
		b := BoxPresenter(box)
		dic = append(dic, b.Map())

	}
	return dic
}

func BoardPrint(b domain.Board) {
	Print(BoardPresenter(b))
}
