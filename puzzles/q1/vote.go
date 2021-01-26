package q1

type Vote struct {
	rock     int
	paper    int
	scissors int
}

func NewVote(r, p, s int) *Vote {
	return &Vote{r, p, s}
}

// 投票完了ではない => どれか2つの値が同じ && 最大値
func (v *Vote) IsDone() bool {
	return !(v.rock == v.paper && v.rock >= v.scissors ||
		v.rock == v.scissors && v.rock >= v.paper ||
		v.paper == v.scissors && v.paper >= v.rock)
}
