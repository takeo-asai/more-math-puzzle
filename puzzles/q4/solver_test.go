package q4_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q4"
)

func Test_q4_Solve(t *testing.T) {
	if q4.Solve(27) != 8800 {
		t.Errorf("q4.solve(27) must be 8800, but it is %d", q4.Solve(27))
	}
	if q4.Solve(30) != 8360 {
		t.Errorf("q4.solve(30) must be 8360, but it is %d", q4.Solve(30))
	}
}
