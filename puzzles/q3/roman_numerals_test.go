package q3_test

import (
	"testing"

	"github.com/takeo-asai/more-math-puzzle/puzzles/q3"
)

func Test_Roman(t *testing.T) {
	if q3.NewRomanNumerals(3888).Roman() != "MMMDCCCLXXXVIII" {
		t.Errorf("3888 is MMMDCCCLXXXVIII, but it is %v", q3.NewRomanNumerals(3888).Roman())
	}
	if q3.NewRomanNumerals(1988).Roman() != "MCMLXXXVIII" {
		t.Errorf("1988 is MCMLXXXVIII, but it is %v", q3.NewRomanNumerals(1988).Roman())
	}
	if q3.NewRomanNumerals(3998).Roman() != "MMMCMXCVIII" {
		t.Errorf("3998 is MMMCMXCVIII, but it is %v", q3.NewRomanNumerals(3998).Roman())
	}
}
