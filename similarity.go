package gnlp

// JaroSimilarity computes Jaro similarity between two sequences.
func JaroSimilarity[T comparable](a, b []T) float64 {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	span := max(len(a), len(b))/2 - 1

	m := 0
	m1 := make([]bool, len(a))
	m2 := make([]bool, len(b))
	for i := 0; i < len(a); i++ {
		for j := max(0, i-span); j < min(len(b), i+span+1); j++ {
			if m2[j] {
				continue
			}
			if a[i] != b[j] {
				continue
			}

			m++
			m1[i] = true
			m2[j] = true
			break
		}
	}

	var t float64 = 0
	var i, j int
	for k := 0; k < m; k++ {
		for !m1[i] {
			i++
		}
		for !m2[j] {
			j++
		}
		if a[i] != b[j] {
			t++
		}
	}
	t /= 2

	return (float64(m)/float64(len(a)) + float64(m)/float64(len(b)) + (float64(m)+t)/float64(m)) / 3.0
}

// JaroWinklerSimilarity computes Jaro-Winkler similarity between two sequences.
// The scaling factor is set to 0.1.
func JaroWinklerSimilarity[T comparable](a, b []T) float64 {
	const p float64 = 0.1

	l := 0
	for ; l < min(4, min(len(a), len(b))); l++ {
		if a[l] != b[l] {
			break
		}
	}

	phi := JaroSimilarity(a, b)
	return phi + float64(l)*p*(1-phi)
}
