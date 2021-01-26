package q2

import (
	"math"
)

// 通過駅はすべて [スタンプを押す、押さない] の2通り
func Solve(source, sink int) int {
	x := math.Abs(float64(source - sink))
	return int(math.Exp2(x-1)+math.Exp2(29-x-1)) - 1
}
