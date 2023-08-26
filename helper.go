package gnlp

import (
	"golang.org/x/exp/constraints"
)

func abs[T constraints.Integer | constraints.Float](x T) T {
	if x >= 0 {
		return x
	}
	return -x
}
