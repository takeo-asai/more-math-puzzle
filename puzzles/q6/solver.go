package q6

func f(x, y int) int {
	var r *Rectangle
	var cnt int
	for r = NewRectangle(x, y); r != nil; r = r.Cut() {
		cnt++
	}
	return cnt
}

func Solve(len, n int) int {
	var cnt int
	for x := 1; x <= len; x++ {
		for y := 1; y <= x; y++ {
			if f(x, y) == n {
				cnt++
			}
		}
	}
	return cnt
}
