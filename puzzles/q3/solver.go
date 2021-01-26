package q3

func Solve(length int) int {
	lengthOfRoman := 0
	for n := 1; n <= 3999; n++ {
		length := NewRomanNumerals(n).Length()
		if length == 12 {
			lengthOfRoman += 1
		}
	}
	return lengthOfRoman
}
