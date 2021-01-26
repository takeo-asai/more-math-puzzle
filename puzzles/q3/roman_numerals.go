package q3

type RomanNumerals struct {
	Number int
}

func NewRomanNumerals(v int) *RomanNumerals {
	return &RomanNumerals{v}
}

func (n *RomanNumerals) Roman() string {
	if n.Number >= 1000 {
		return "M" + NewRomanNumerals(n.Number-1000).Roman()
	}
	if n.Number >= 900 {
		return "CM" + NewRomanNumerals(n.Number-900).Roman()
	}
	if n.Number >= 500 {
		return "D" + NewRomanNumerals(n.Number-500).Roman()
	}
	if n.Number >= 400 {
		return "CD" + NewRomanNumerals(n.Number-400).Roman()
	}
	if n.Number >= 100 {
		return "C" + NewRomanNumerals(n.Number-100).Roman()
	}
	if n.Number >= 90 {
		return "XC" + NewRomanNumerals(n.Number-90).Roman()
	}
	if n.Number >= 50 {
		return "L" + NewRomanNumerals(n.Number-50).Roman()
	}
	if n.Number >= 40 {
		return "XL" + NewRomanNumerals(n.Number-40).Roman()
	}
	if n.Number >= 10 {
		return "X" + NewRomanNumerals(n.Number-10).Roman()
	}
	if n.Number >= 9 {
		return "IX" + NewRomanNumerals(n.Number-9).Roman()
	}
	if n.Number >= 5 {
		return "V" + NewRomanNumerals(n.Number-5).Roman()
	}
	if n.Number >= 4 {
		return "IV" + NewRomanNumerals(n.Number-4).Roman()
	}
	if n.Number >= 1 {
		return "I" + NewRomanNumerals(n.Number-1).Roman()
	}
	return ""
}

func (n *RomanNumerals) Length() int {
	return len(n.Roman())
}
