package q0

func f(m *Memo, pre int, remain int) int {
	if remain < 0 {
		return 0
	}
	if remain == 0 {
		return 1
	}
	var sum int
	for k := pre; k <= 10 && (remain-k) >= 0; k++ {
		if v, ok := m.Get(k, remain-k); ok {
			sum += v
		} else {
			v := f(m, k, remain-k)
			m.Set(k, remain-k, v)
			sum += v
		}
	}
	return sum
}

func Solve(n int) int {
	m := NewMemo()
	return f(m, 2, n)
}
