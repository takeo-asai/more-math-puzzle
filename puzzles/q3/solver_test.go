package q3_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q3"
)

func Test_q3_Solve(t *testing.T) {
	if q3.Solve(12) != 93 {
		t.Errorf("q3.solve(12) must be 93, but it is %d", q3.Solve(12))
	}
}
