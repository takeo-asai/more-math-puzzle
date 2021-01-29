package q4

type Number struct {
	value int
}

func NewNumber(k int) *Number {
	if k < 0 || k > 9 {
		return nil
	}
	return &Number{k}
}

func (n *Number) Lights() int {
	switch n.value {
	case 0:
		return 6
	case 1:
		return 2
	case 2:
		return 5
	case 3:
		return 5
	case 4:
		return 4
	case 5:
		return 5
	case 6:
		return 6
	case 7:
		return 3
	case 8:
		return 7
	case 9:
		return 6
	default:
		return 9223372036854775807
	}
}
