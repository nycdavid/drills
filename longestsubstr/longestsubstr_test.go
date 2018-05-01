package longestsubstr

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
	}
	for substring, lgth := range cases {
		res := lengthOfLongestSubstring(substring)

		if res != lgth {
			msg := fmt.Sprintf("Expected substring %s to be %d, got %d", substring, lgth, res)
			t.Error(msg)
		}
	}
}
