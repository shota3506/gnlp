package gnlp_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/shota3506/gnlp"
)

func TestJaccardIndex(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected float64
	}{
		{[]rune(""), []rune(""), 1},
		{[]rune("a"), []rune(""), 0},
		{[]rune("a"), []rune("a"), 1},
		{[]rune("ab"), []rune("a"), 0.5},
		{[]rune("abc"), []rune("aa"), 1.0 / 3.0},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d := gnlp.JaccardIndex(tc.a, tc.b)
			if math.Abs(tc.expected-d) > 1e-9 {
				t.Errorf("epected: %f, actual: %f", tc.expected, d)
			}
		})
	}
}

func TestDiceIndex(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected float64
	}{
		{[]rune(""), []rune(""), 1},
		{[]rune("a"), []rune(""), 0},
		{[]rune("a"), []rune("a"), 1},
		{[]rune("ab"), []rune("a"), 2.0 / 3.0},
		{[]rune("aba"), []rune("ac"), 0.5},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d := gnlp.DiceIndex(tc.a, tc.b)
			if math.Abs(tc.expected-d) > 1e-9 {
				t.Errorf("epected: %f, actual: %f", tc.expected, d)
			}
		})
	}
}

func TestSimpsonIndex(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected float64
	}{
		{[]rune(""), []rune(""), 1},
		{[]rune("a"), []rune(""), 1},
		{[]rune("a"), []rune("a"), 1},
		{[]rune("ab"), []rune("a"), 1},
		{[]rune("aba"), []rune("ac"), 0.5},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d := gnlp.SimpsonIndex(tc.a, tc.b)
			if math.Abs(tc.expected-d) > 1e-9 {
				t.Errorf("epected: %f, actual: %f", tc.expected, d)
			}
		})
	}
}
