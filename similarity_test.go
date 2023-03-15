package gnlp_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/shota3506/gnlp"
)

func TestJaroSimilarity(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected float64
	}{
		{[]rune(""), []rune(""), 0},
		{[]rune("a"), []rune(""), 0},
		{[]rune(""), []rune("a"), 0},
		{[]rune("CRATE"), []rune("TRACE"), 0.7333333333},
		{[]rune("DwAyNE"), []rune("DuANE"), 0.8222222222},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d := gnlp.JaroSimilarity(tc.a, tc.b)
			if math.Abs(tc.expected-d) > 1e-9 {
				t.Errorf("epected: %f, actual: %f", tc.expected, d)
			}
		})
	}
}

func TestJaroWinklerSimilarity(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected float64
	}{
		{[]rune(""), []rune(""), 0},
		{[]rune("a"), []rune(""), 0},
		{[]rune(""), []rune("a"), 0},
		{[]rune("TRATE"), []rune("TRACE"), 0.9066666667},
		{[]rune("DwAyNE"), []rune("DuANE"), 0.84},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d := gnlp.JaroWinklerSimilarity(tc.a, tc.b)
			if math.Abs(tc.expected-d) > 1e-9 {
				t.Errorf("epected: %f, actual: %f", tc.expected, d)
			}
		})
	}
}
