package gnlp

import (
	"errors"
)

// HammingDistance computes Hamming distance between two sequences
// of the same length.
func HammingDistance[T comparable](a, b []T) (int64, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences must be the same length")
	}

	var d int64
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}

// LevenshteinDistance computes Levenshtein distance between two sequences.
func LevenshteinDistance[T comparable](a, b []T) int64 {
	d := make([]int64, len(a)+1)
	for i := 0; i <= len(a); i++ {
		d[i] = int64(i)
	}

	for j := 1; j <= len(b); j++ {
		d[0] = int64(j)
		last := int64(j - 1)
		for i := 1; i <= len(a); i++ {
			prev := d[i]

			var cost int64 = 0
			if a[i-1] != b[j-1] {
				cost = 1
			}

			d1 := d[i] + 1
			d2 := d[i-1] + 1
			d3 := last + cost

			d[i] = min(min(d1, d2), d3)

			last = prev
		}
	}

	return d[len(a)]
}

// DamerauLevenshteinDistance computes Damerau-Levenshtein distance between two sequences.
func DamerauLevenshteinDistance[T comparable](a, b []T) int64 {
	d := make([][]int64, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		d[i] = make([]int64, len(b)+1)
	}

	for j := 1; j <= len(b); j++ {
		d[0][j] = int64(j)
	}
	for i := 1; i <= len(a); i++ {
		d[i][0] = int64(i)
		for j := 1; j <= len(b); j++ {
			var cost int64 = 0
			if a[i-1] != b[j-1] {
				cost = 1
			}

			d1 := d[i][j-1] + 1
			d2 := d[i-1][j] + 1
			d3 := d[i-1][j-1] + cost

			d[i][j] = min(min(d1, d2), d3)

			if i >= 2 && j >= 2 && a[i-1] == b[j-2] && a[i-2] == b[j-1] {
				d[i][j] = min(d[i][j], d[i-2][j-2]+cost)
			}
		}
	}
	return d[len(a)][len(b)]
}
