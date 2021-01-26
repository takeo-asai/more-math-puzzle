package q1_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q1"
)

func Test_q1_Solver(t *testing.T) {
	if q1.Solve(2) != 1 {
		t.Errorf("q1.Solve(2) must be 1, but it is %d", q1.Solve(2))
	}
	if q1.Solve(3) != 1 {
		t.Errorf("q1.Solve(3) must be 1, but it is %d", q1.Solve(3))
	}
	if q1.Solve(4) != 2 {
		t.Errorf("q1.Solve(4) must be 2, but it is %d", q1.Solve(4))
	}
	if q1.Solve(6) != 4 {
		t.Errorf("q1.Solve(6) must be 4, but it is %d", q1.Solve(6))
	}
	if q1.Solve(100) != 437420 {
		t.Errorf("q1.Solve(100) must be 437420, but it is %d", q1.Solve(100))
	}
}
