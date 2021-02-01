package q7

func Math_nPr(n, r int) int {
	var p int = 1
	for i := 0; i < r; i++ {
		p *= (n - i)
	}
	return p
}
