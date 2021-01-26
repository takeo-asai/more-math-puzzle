package q1

func Solve(n int) int {
	nVotes := 0
	for r := 0; r <= n; r++ {
		for p := 0; p <= n-r; p++ {
			s := n - r - p
			if NewVote(r, p, s).IsDone() {
				nVotes += 1
			}
		}
	}
	return nVotes
}
