package gnlp

import (
	"math"
)

// JaccardIndex computes Jaccard index (Jaccard similarity coefficient)
// of two sets.
func JaccardIndex[T comparable](a, b []T) float64 {
	if len(a) == 0 && len(b) == 0 {
		return 1
	}

	// union
	union := map[T]any{}
	for _, x := range a {
		union[x] = struct{}{}
	}
	for _, x := range b {
		union[x] = struct{}{}
	}

	// intersection
	mapping := map[T]any{}
	for _, x := range a {
		mapping[x] = struct{}{}
	}
	intersection := map[T]any{}
	for _, x := range b {
		if _, ok := mapping[x]; ok {
			intersection[x] = struct{}{}
		}
	}

	return float64(len(intersection)) / float64(len(union))
}

// DiceIndex computes Sørensen-Dice index (Sørensen-Dice similarity coefficient)
// of two sets.
func DiceIndex[T comparable](a, b []T) float64 {
	if len(a) == 0 && len(b) == 0 {
		return 1
	}

	aMap := map[T]any{}
	for _, x := range a {
		aMap[x] = struct{}{}
	}
	bMap := map[T]any{}
	for _, x := range b {
		bMap[x] = struct{}{}
	}
	intersection := map[T]any{}
	for _, x := range b {
		if _, ok := aMap[x]; ok {
			intersection[x] = struct{}{}
		}
	}

	return 2 * float64(len(intersection)) / float64(len(aMap)+len(bMap))
}

// SimpsonIndex computes Szymkiewicz–Simpson index (Szymkiewicz–Simpson similarity coefficient)
// of two sets.
func SimpsonIndex[T comparable](a, b []T) float64 {
	if len(a) == 0 || len(b) == 0 {
		return 1
	}

	aMap := map[T]any{}
	for _, x := range a {
		aMap[x] = struct{}{}
	}
	bMap := map[T]any{}
	for _, x := range b {
		bMap[x] = struct{}{}
	}
	intersection := map[T]any{}
	for _, x := range b {
		if _, ok := aMap[x]; ok {
			intersection[x] = struct{}{}
		}
	}

	return float64(len(intersection)) / math.Min(float64(len(aMap)), float64(len(bMap)))
}
