package main

type Samurai struct {
	Owner   Owner
	Name    string
	Ability Ability
}

var samurai struct {
	Deck []*Samurai
}

func init() {
	samurai.Deck = []*Samurai{
		NewSamurai("弥助", 1, 2, 6, CHANGE_DICE),
		NewSamurai("木下藤吉郎", 3, 4, 1, CHANGE_DICE),
		NewSamurai("明智光秀", 5, 6, 1, CHANGE_DICE),
		NewSamurai("柴田勝家", 3, 4, 5, CHANGE_DICE),
		NewSamurai("前田利家", 5, 6, 3, CHANGE_DICE),
		NewSamurai("滝川一益", 1, 2, 4, CHANGE_DICE),
		NewSamurai("毛利元就", 5, 6, 4, CHANGE_DICE),
		NewSamurai("徳川家康", 3, 4, 2, CHANGE_DICE),
		NewSamurai("今川義元", 5, 6, 1, CHANGE_DICE),
		NewSamurai("北条氏康", 1, 2, 3, CHANGE_DICE),
		NewSamurai("島津義弘", 3, 4, 6, CHANGE_DICE),
		NewSamurai("松永久秀", 0, 0, 0, CHANGE_SAMURAI),
		NewSamurai("武田信玄", 1, 0, 0, ADD_GOLD),
		NewSamurai("上杉謙信", 2, 0, 0, ADD_BATTLE_DICE),
		NewSamurai("伊達政宗", 2, 4, 0, CHANGE_AREA),
		NewSamurai("九鬼義隆", 2, 6, 0, CHANGE_AREA),
		NewSamurai("伊藤マンショ", 1, 6, 0, CHANGE_AREA),
		NewSamurai("足利義昭", 1, 0, 0, DELETE_DICE),
		NewSamurai("大友宗麟", 1, 4, 0, CHANGE_AREA),
	}
}

func NewSamurai(name string, x, y, z int, a AbilityType) *Samurai {
	return &Samurai{
		Name:    name,
		Ability: NewAbility(a, x, y, z),
	}
}

func (s *Samurai) Map() map[string]string {
	return map[string]string{
		"name":    s.Name,
		"ability": s.Ability.String(),
	}
}
