package q7_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q7"
)

func Test_q7_Solver(t *testing.T) {
	if q7.Solve(2) != 1 {
		t.Errorf("q7.Solve(2) == %d", q7.Solve(2))
	}
	if q7.Solve(3) != 8 {
		t.Errorf("q7.Solve(3) == %d", q7.Solve(3))
	}
	if q7.Solve(15) != 17368162415924 {
		t.Errorf("q7.Solve(15) == %d", q7.Solve(15))
	}
}
