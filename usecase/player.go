package usecase

import (
	"github.com/tamura2004/nobugo/domain"
	"log"
)

type Player struct {
	domain.Player
}

// プレイヤーに武将カードを１枚配る
func (p *Player) DrawSamurai(g *Game) {
	deck := g.Samurai
	card := deck.Draw()
	if card.Type == domain.EMPTY {
		log.Fatal("samurai deck is empty")
	}
	p.Samurai.Bottom(card)
}
