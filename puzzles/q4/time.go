package q4

type Time struct {
	h0 *Number
	h1 *Number
	m0 *Number
	m1 *Number
	s0 *Number
	s1 *Number
}

func NewTime(h int, m int, s int) *Time {
	if h < 0 || h > 23 || m < 0 || m > 59 || s < 0 || s > 59 {
		return nil
	}
	return &Time{NewNumber(h / 10), NewNumber(h % 10), NewNumber(m / 10), NewNumber(m % 10), NewNumber(s / 10), NewNumber(s % 10)}
}

func (t *Time) Lights() int {
	return t.h0.Lights() + t.h1.Lights() + t.m0.Lights() + t.m1.Lights() + t.s0.Lights() + t.s1.Lights()
}
