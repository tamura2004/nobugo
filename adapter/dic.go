package adapter

type Dic []map[string][]string

func (d Dic) Width() int {
	return len(d)
}

func (d Dic) Height(key string) int {
	height := 0
	for _, dic := range d {
		l := len(dic[key])
		if height < l {
			height = l
		}
	}
	return height
}
