package q0_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q0"
)

func Test_memo(t *testing.T) {
	m := q0.NewMemo()
	m.Set(0, 0, 100)
	if v, _ := m.Get(0, 0); v != 100 {
		t.Errorf("%d must be 100", v)
	}
	if _, b := m.Get(0, 1); b {
		t.Errorf("m.Get(1) must not exist")
	}
}
