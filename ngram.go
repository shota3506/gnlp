package gnlp

// NGrams returns a contiguous sequence of n items
// from the given sequence.
func NGrams[T any](seq []T, n int) (ngram [][]T) {
	if n == 0 {
		return
	}
	for i := 0; i+n <= len(seq); i++ {
		ngram = append(ngram, seq[i:i+n])
	}
	return ngram
}
