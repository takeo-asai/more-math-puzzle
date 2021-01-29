package q4_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q4"
)

func Test_q4_Time_Lights(t *testing.T) {
	if q4.NewTime(12, 34, 56).Lights() != 27 {
		t.Errorf("q4.Time(12, 34, 56) must be 27, but it is %d", q4.NewTime(12, 34, 56).Lights())
	}
}
