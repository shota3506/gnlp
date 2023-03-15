package gnlp

import (
	"golang.org/x/exp/constraints"
)

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func abs[T constraints.Integer | constraints.Float](x T) T {
	if x >= 0 {
		return x
	}
	return -x
}
