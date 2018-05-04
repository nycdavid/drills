package longestsubstr

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := map[string]int{
		"c":        1,
		"bbbbb":    1,
		"aab":      2,
		"abcabcbb": 3,
		"pwwkew":   3,
		"dvdf":     3,
	}

	for substring, lgth := range cases {
		res := lengthOfLongestSubstring(substring)

		if res != lgth {
			msg := fmt.Sprintf("Expected substring %s to be %d, got %d", substring, lgth, res)
			t.Error(msg)
		}
	}
}

func TestContinueAfterDuplicateChar(t *testing.T) {
	cases := map[string]int{
		"anviaj":          5,
		"anviajbc":        7,
		"anviajbcapolice": 6,
	}
	for substring, lgth := range cases {
		res := lengthOfLongestSubstring(substring)

		if res != lgth {
			msg := fmt.Sprintf("Expected substring %s to be %d, got %d", substring, lgth, res)
			t.Error(msg)
		}
	}
}
