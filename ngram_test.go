package gnlp_test

import (
	"testing"

	"github.com/shota3506/gnlp"
)

func TestNGram(t *testing.T) {
	ngram := gnlp.NGrams([]rune("abcdef"), 2)
	expected := [][]rune{{'a', 'b'}, {'b', 'c'}, {'c', 'd'}, {'d', 'e'}, {'e', 'f'}}
	if len(expected) != len(ngram) {
		t.Fatalf("expected length: %d, actual length: %d", len(expected), len(ngram))
	}
	for i := 0; i < len(expected); i++ {
		if len(expected[i]) != len(ngram[i]) {
			t.Fatalf("expected length: %d, actual length: %d", len(expected[i]), len(ngram[i]))
		}
		for j := 0; j < len(expected[i]); j++ {
			if expected[i][j] != ngram[i][j] {
				t.Fatalf("expected: %d, actual: %d", expected[i][j], ngram[i][j])
			}
		}
	}
}
