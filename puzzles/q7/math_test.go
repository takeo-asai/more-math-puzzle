package q7_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q7"
)

func Test_q7_Math_nPr(t *testing.T) {
	if q7.Math_nPr(3, 0) != 1 {
		t.Errorf("q7.Math_nPr(3, 0) == %d", q7.Math_nPr(3, 0))
	}
	if q7.Math_nPr(4, 3) != 24 {
		t.Errorf("q7.Math_nPr(4, 3) == %d", q7.Math_nPr(4, 3))
	}
	if q7.Math_nPr(5, 2) != 20 {
		t.Errorf("q7.Math_nPr(5, 2) == %d", q7.Math_nPr(5, 2))
	}
}
