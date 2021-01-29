package q5

func Solve(n int) int {
	var cnt int
	for _, v := range PascalTriangle(n) {
		cnt += Coins(v)
	}
	return cnt
}
