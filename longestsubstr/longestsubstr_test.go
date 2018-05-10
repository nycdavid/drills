package longestsubstr

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := map[string]int{
		"c":               1,
		"bbbbb":           1,
		"aab":             2,
		"abcabcbb":        3,
		"pwwkew":          3,
		"dvdf":            3,
		"anviaj":          5,
		"anviajbcapolice": 8,
		"bb":              1,
		"ggububgvfk":      6,
	}

	for substring, lgth := range cases {
		res := lengthOfLongestSubstring(substring)

		if res != lgth {
			msg := fmt.Sprintf("Expected substring %s to be %d, got %d", substring, lgth, res)
			t.Error(msg)
		}
	}
}
