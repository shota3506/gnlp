package gnlp

import (
	"math"
)

// BLEU computes a sentence-level BLEU score.
//
// Papineni, Kishore, Salim Roukos, Todd Ward, and Wei-Jing Zhu. 2002.
// "BLEU: a method for automatic evaluation of machine translation."
// In Proceedings of ACL. https://www.aclweb.org/anthology/P02-1040.pdf
//
// The candidate parameter is a sequence of token and the references parameter is
// a set of sequences of token.
// This method returns zero if there's no refernece.
func BLEU[T comparable](candidate []T, references [][]T) float64 {
	return CorpusBLEU([][]T{candidate}, [][][]T{references})
}

// CorpusBLEU computes a corpus-level BLEU score.
//
// Papineni, Kishore, Salim Roukos, Todd Ward, and Wei-Jing Zhu. 2002.
// "BLEU: a method for automatic evaluation of machine translation."
// In Proceedings of ACL. https://www.aclweb.org/anthology/P02-1040.pdf
//
// The candidate list and references list should be the same length.
// Otherwise it returns zero.
//
// Note that this method doesn't return the average of sentence-level BLEU score.
// It calculates the micro-average of precision as the original BLEU paper.
func CorpusBLEU[T comparable](candidateList [][]T, referencesList [][][]T) float64 {
	if len(candidateList) != len(referencesList) {
		// candidate list and references list should be the same length
		return 0
	}

	numerators := make([]int64, 4+1)
	denominators := make([]int64, 4+1)
	totalLength := 0
	totalRefLength := 0
	for idx := 0; idx < len(candidateList); idx++ {
		candidate := candidateList[idx]
		references := referencesList[idx]
		if len(references) == 0 {
			continue
		}

		for n := 1; n <= 4; n++ {
			nn, dn := bleuModifiedPrecision(candidate, references, n)
			numerators[n] += nn // no smoothing
			denominators[n] += dn
		}

		length := len(candidate)
		refLength := len(references[0])
		for _, refernece := range references[1:] {
			rlength := len(refernece)
			if abs(length-rlength) < abs(length-refLength) ||
				(abs(length-rlength) == abs(length-refLength) && rlength < refLength) {
				refLength = rlength
			}
		}
		totalLength += length
		totalRefLength += refLength
	}

	var numerator, denominator int64 = 1, 1
	for n := 1; n <= 4; n++ {
		numerator *= numerators[n]
		denominator *= denominators[n]
	}
	if numerator == 0 || denominator == 0 {
		return 0
	}

	var precision float64 = float64(numerator) / float64(denominator)

	bp := bleuBrevityPenalty(totalLength, totalRefLength)

	return bp * math.Pow(precision, 0.25)
}

func bleuModifiedPrecision[T comparable](candidate []T, references [][]T, n int) (int64, int64) {
	c := NGrams(candidate, n)
	if len(c) == 0 {
		return 0, 0
	}

	helper := func(a, b []T) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	match := make([]bool, len(c))
	for _, reference := range references {
		r := NGrams(reference, n)
		refMatch := make([]bool, len(r))
		for i := 0; i < len(c); i++ {
			for j := 0; j < len(r); j++ {
				if refMatch[j] {
					continue
				}
				if !helper(c[i], r[j]) {
					continue
				}
				match[i] = true
				refMatch[j] = true
				break
			}
		}
	}

	var numerator int64 = 0
	for _, b := range match {
		if b {
			numerator++
		}
	}
	return numerator, int64(len(c))
}

func bleuBrevityPenalty(length, refLength int) float64 {
	if length == 0 {
		return 0 // avoid zero division
	}
	if length >= refLength {
		return 1
	}
	return math.Exp(1 - float64(refLength)/float64(length))
}

// ROUGEN computes a ROUGE-N score,
// which is a recall-oriented text summarization metrics.
//
// Chin-Yew Lin. 2004.
// "ROUGE: A Package for Automatic Evaluation of Summaries."
// In Proceedings of ACL. https://aclanthology.org/W04-1013.pdf
func ROUGEN[T comparable](candidate []T, references [][]T, n int) float64 {
	c := NGrams(candidate, n)

	helper := func(a, b []T) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	var numerator, denominator int
	for _, reference := range references {
		r := NGrams(reference, n)
		if len(r) == 0 {
			continue
		}
		match := make([]bool, len(c))
		for j := 0; j < len(r); j++ {
			for i := 0; i < len(c); i++ {
				if match[i] {
					continue
				}
				if !helper(c[i], r[j]) {
					continue
				}
				match[i] = true
				numerator++
			}
		}
		denominator += len(r)
	}

	if denominator == 0 {
		return 0
	}
	return float64(numerator) / float64(denominator)
}
