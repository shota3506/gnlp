package gnlp

import (
	"golang.org/x/exp/slices"
)

// LongestCommonSubsequences returns longest subsequences commmon
// to given two sequences.
// It returns all valid subsequeces.
//
// This method returns a slice which contains at least one sequence.
// It returns [][]T{{}} if there's no common subsequence.
func LongestCommonSubsequences[T comparable](a, b []T) [][]T {
	d := make([][]int64, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		d[i] = make([]int64, len(b)+1)
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				d[i][j] = d[i-1][j-1] + 1
			} else {
				d[i][j] = max(d[i-1][j], d[i][j-1])
			}
		}
	}

	var lcss [][]T
	var trackBackFunc func([]T, int, int)
	trackBackFunc = func(s []T, i, j int) {
		if i == 0 || j == 0 {
			if len(s) > 0 {
				lcss = append(lcss, s)
			}
			return
		}

		if a[i-1] == b[j-1] {
			q := make([]T, 1, len(s)+1)
			q[0] = a[i-1]
			q = append(q, s...)
			trackBackFunc(q, i-1, j-1)
		} else {
			if d[i-1][j] == d[i][j] {
				trackBackFunc(s, i-1, j)
			}
			if d[i][j-1] == d[i][j] {
				trackBackFunc(s, i, j-1)
			}
		}
	}

	trackBackFunc([]T{}, len(a), len(b))

	if len(lcss) == 0 {
		// empty sequence
		return [][]T{{}}
	}
	return uniqueSequences(lcss)
}

func uniqueSequences[T comparable](ss [][]T) [][]T {
	type node map[T]node

	root := make(map[T]node)
	for _, s := range ss {
		n := root
		for _, t := range s {
			if _, ok := n[t]; !ok {
				n[t] = make(map[T]node)
			}
			n = n[t]
		}
	}

	var uniq [][]T
	var traverse func(n node, s []T)
	traverse = func(n node, s []T) {
		if len(n) == 0 {
			if len(s) > 0 {
				uniq = append(uniq, s)
			}
			return
		}

		for t, next := range n {
			traverse(next, append(slices.Clone(s), t))
		}
	}

	traverse(root, []T{})

	return uniq
}
