package samurai

import (
	"github.com/tamura2004/nobugo/samurai/ability"
)

type Samurai struct {
	Name    string
	Ability *ability.Ability
}

func New(name string, x, y, z, a int) Samurai {
	return Samurai{
		Name:    name,
		Ability: ability.New(a, x, y, z),
	}
}

func (s Samurai) Value() []string {
	return []string{
		s.Name,
		s.Ability.String(),
	}
}
