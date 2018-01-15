package counter

type Counter map[interface{}]int

func New() Counter {
	return make(Counter)
}

func (c Counter) Add(key interface{}, num int) {
	c[key] += num
}

func (c Counter) Max() (m int) {
	for _, v := range c {
		if m < v {
			m = v
		}
	}
	return
}
