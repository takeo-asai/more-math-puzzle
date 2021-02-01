package q7

func Solve(n int) int {
	var cnt int = 0
	for i := 1; i < n; i++ {
		cnt += i * (n - i) * Math_nPr(n, i-1)
	}
	return cnt
}
