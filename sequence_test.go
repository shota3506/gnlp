package gnlp_test

import (
	"fmt"
	"github.com/shota3506/gnlp"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	for i, tc := range []struct {
		a        []rune
		b        []rune
		expected [][]rune
	}{
		{[]rune(""), []rune("ABC"), [][]rune{{}}},
		{[]rune("ABCDBCD"), []rune("BCDBAC"), [][]rune{[]rune("BCDBC")}},
		{[]rune("AGCAT"), []rune("GAC"), [][]rune{[]rune("AC"), []rune("GC"), []rune("GA")}},
		{[]rune("ABCDBCDA"), []rune("BDCABCDBA"), [][]rune{[]rune("ABCDBA"), []rune("BCBCDA"), []rune("BDBCDA")}},
	} {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			lcss := gnlp.LongestCommonSubsequences(tc.a, tc.b)

			if len(lcss) != len(tc.expected) {
				t.Fatalf("expected number of lcs: %d, actual number of lcs: %d", len(tc.expected), len(lcss))
			}

			memo := make(map[string]any)
			for _, e := range tc.expected {
				memo[string(e)] = struct{}{}
			}
			match := true
			for _, lcs := range lcss {
				key := string(lcs)
				if _, ok := memo[key]; !ok {
					match = false
					break
				}
				delete(memo, key)
			}
			if !match {
				t.Errorf("expected: %v, actual: %v", tc.expected, lcss)
			}
		})
	}
}
