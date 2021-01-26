package q1

type Memo struct {
	values map[int]map[int]int
}

func NewMemo() *Memo {
	return &Memo{make(map[int]map[int]int)}
}

func (m *Memo) Set(i int, j int, value int) *Memo {
	if _, ok := m.values[i]; ok {
		m.values[i][j] = value
	} else {
		m.values[i] = make(map[int]int)
		m.values[i][j] = value
	}
	return m
}

func (m *Memo) Get(i int, j int) (int, bool) {
	v, ok := m.values[i][j]
	return v, ok
}
