package q1_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q1"
)

func Test_q1_Solver(t *testing.T) {
	if q1.Solve(4) != 12 {
		t.Errorf("q0.Solve(4) must be 12, but it is %d", q1.Solve(4))
	}
	if q1.Solve(100) != 5100 {
		t.Errorf("q1.Solve(100) must be 5100, but it is %d", q1.Solve(100))
	}

}
