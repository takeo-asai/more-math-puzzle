package q2_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q2"
)

func Test_q2_Solve(t *testing.T) {
	if q2.Solve(1, 17) != 36863 {
		t.Errorf("q2.Solve(1,17) must be 36863, but it is %d", q2.Solve(1, 17))
	}
}
