package q5_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q5"
)

func Test_q4_Solve(t *testing.T) {
	if q5.Solve(0) != 1 {
		t.Errorf("q5.solve(0) must be 1, but it is %d", q5.Solve(0))
	}
	if q5.Solve(1) != 2 {
		t.Errorf("q5.solve(1) must be 2, but it is %d", q5.Solve(1))
	}
	if q5.Solve(2) != 4 {
		t.Errorf("q5.solve(2) must be 4, but it is %d", q5.Solve(2))
	}
	if q5.Solve(3) != 8 {
		t.Errorf("q5.solve(3) must be 8, but it is %d", q5.Solve(3))
	}
	if q5.Solve(4) != 12 {
		t.Errorf("q5.solve(4) must be 12, but it is %d", q5.Solve(4))
	}
	if q5.Solve(5) != 6 {
		t.Errorf("q5.solve(5) must be 6, but it is %d", q5.Solve(5))
	}
	if q5.Solve(6) != 12 {
		t.Errorf("q5.solve(6) must be 12, but it is %d", q5.Solve(6))
	}
	if q5.Solve(7) != 22 {
		t.Errorf("q5.solve(7) must be 22, but it is %d", q5.Solve(7))
	}
	if q5.Solve(8) != 31 {
		t.Errorf("q5.solve(8) must be 31, but it is %d", q5.Solve(8))
	}
	if q5.Solve(9) != 48 {
		t.Errorf("q5.solve(9) must be 48, but it is %d", q5.Solve(9))
	}
	if q5.Solve(45) != 3518437540 {
		t.Errorf("q5.solve(45) must be 3518437540, but it is %d", q5.Solve(45))
	}
}
