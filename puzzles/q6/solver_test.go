package q6_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q6"
)

func Test_q6_Solve(t *testing.T) {
	if q6.Solve(8, 5) != 8 {
		t.Errorf("q6.Solve(8, 5) == %d", q6.Solve(8, 5))
	}
	if q6.Solve(1000, 20) != 26882 {
		t.Errorf("q6.Solve(1000, 20) == %d", q6.Solve(1000, 20))
	}
}
