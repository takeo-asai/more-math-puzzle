package q5_test

import (
	"reflect"
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q5"
)

func Test_q5_Triangle(t *testing.T) {
	if !reflect.DeepEqual(q5.PascalTriangle(5), []int{1, 5, 10, 10, 5, 1}) {
		t.Errorf("q5.PascalTriangle(5) is %v", q5.PascalTriangle(5))
	}
}
