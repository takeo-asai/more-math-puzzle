package q4

func Solve(n int) int {
	var cnt int
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {
			for s := 0; s < 60; s++ {
				if NewTime(h, m, s).Lights() == n {
					cnt += 1
				}
			}
		}
	}
	return cnt
}
