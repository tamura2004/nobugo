package castle

import (
	"fmt"
	"github.com/tamura2004/nobugo/castle/area"
)

type Castle struct {
	Area    area.Area
	Country string
	Name    string
	Defense Defense
}

func New(country, name string, x, y, z int, a area.Area) Castle {
	return Castle{
		Name:    name,
		Country: country,
		Area:    a,
		Defense: [3]int{x, y, z},
	}
}

func (c Castle) Value() []string {
	return []string{
		c.Name,
		c.Defense.String(),
		fmt.Sprintf("%s/%s", c.Area, c.Country),
	}
}

type Defense [3]int

func (d *Defense) String() (s string) {
	for _, v := range *d {
		if v != 0 {
			s += fmt.Sprintf("[%d]", v)
		}
	}
	return
}
