package gnlp_test

import (
	"fmt"
	"testing"

	"github.com/shota3506/gnlp"
)

func TestHammingDistance(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected int64
	}{
		{[]rune("abcde"), []rune("abcde"), 0},
		{[]rune("abcde"), []rune("abxxe"), 2},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d, err := gnlp.HammingDistance(tc.a, tc.b)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tc.expected != d {
				t.Errorf("epected: %d, actual: %d", tc.expected, d)
			}
		})
	}
}

func TestHammingDistanceError(t *testing.T) {
	for i, tc := range []struct {
		a []rune
		b []rune
	}{
		{[]rune("a"), []rune("")},
		{[]rune("abc"), []rune("ab")},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			_, err := gnlp.HammingDistance(tc.a, tc.b)
			if err == nil {
				t.Error("error epected")
			}
		})
	}
}

func TestLevenshteinDistance(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected int64
	}{
		{[]rune(""), []rune(""), 0},
		{[]rune("abcde"), []rune("abcde"), 0},
		{[]rune("a"), []rune(""), 1},
		{[]rune("abcde"), []rune("ce"), 3},
		{[]rune("ce"), []rune("abcde"), 3},
		{[]rune("abcde"), []rune("ed"), 4},
		{[]rune("abcde"), []rune("acbde"), 2},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d := gnlp.LevenshteinDistance(tc.a, tc.b)
			if tc.expected != d {
				t.Errorf("epected: %d, actual: %d", tc.expected, d)
			}
		})
	}
}

func TestDamerauLevenshteinDistance(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected int64
	}{
		{[]rune(""), []rune(""), 0},
		{[]rune("abcde"), []rune("abcde"), 0},
		{[]rune("a"), []rune(""), 1},
		{[]rune("abcde"), []rune("ce"), 3},
		{[]rune("ce"), []rune("abcde"), 3},
		{[]rune("abcde"), []rune("ed"), 4},
		{[]rune("abcde"), []rune("acbde"), 1},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			d := gnlp.DamerauLevenshteinDistance(tc.a, tc.b)
			if tc.expected != d {
				t.Errorf("epected: %d, actual: %d", tc.expected, d)
			}
		})
	}
}
