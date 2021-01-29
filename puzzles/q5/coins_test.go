package q5_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q5"
)

func Test_q5_Coins(t *testing.T) {
	if q5.Coins(6) != 2 {
		t.Errorf("q5.Coins(6) == 2, but %d", q5.Coins(6))
	}
}
