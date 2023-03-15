package gnlp_test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/shota3506/gnlp"
)

func TestBLEU(t *testing.T) {
	candidate := strings.Split("The NASA Opportunity rover is battling a massive dust storm on Mars .", " ")
	references := [][]string{
		strings.Split("The Opportunity rover is combating a big sandstorm on Mars .", " "),
		strings.Split("A NASA rover is fighting a massive storm on Mars .", " "),
	}

	bleu := gnlp.BLEU(candidate, references)
	expected := 0.3277456805
	if math.Abs(expected-bleu) > 1e-9 {
		t.Errorf("epected: %f, actual: %f", expected, bleu)
	}
}

func TestCorpusBLEU(t *testing.T) {
	for i, tc := range []struct {
		candidateList  [][]string
		referencesList [][][]string
		expected       float64
	}{
		{
			candidateList: [][]string{
				strings.Split("The NASA Opportunity rover is battling a massive dust storm on Mars .", " "),
			},
			referencesList: [][][]string{
				{
					strings.Split("The Opportunity rover is combating a big sandstorm on Mars .", " "),
					strings.Split("A NASA rover is fighting a massive storm on Mars .", " "),
				},
			},
			expected: 0.3277456805,
		},
		{
			candidateList: [][]string{
				strings.Split("It is a guide to action which ensures that the military always obeys the commands of the party", " "),
				strings.Split("he read the book because he was interested in world history", " "),
			},
			referencesList: [][][]string{
				{
					strings.Split("It is a guide to action that ensures that the military will forever heed Party commands", " "),
					strings.Split("It is the guiding principle which guarantees the military forces always being under the command of the Party", " "),
					strings.Split("It is the practical guide for the army always to heed the directions of the party", " "),
				},
				{
					strings.Split("he was interested in world history because he read the book", " "),
				},
			},
			expected: 0.5920778869,
		},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			bleu := gnlp.CorpusBLEU(tc.candidateList, tc.referencesList)
			if math.Abs(tc.expected-bleu) > 1e-9 {
				t.Errorf("epected: %f, actual: %f", tc.expected, bleu)
			}
		})
	}
}
