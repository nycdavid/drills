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
		res, _ := lngthOfLongestSubstring(substring)

		if res != lgth {
			msg := fmt.Sprintf("Expected substring %s to be %d, got %d", substring, lgth, res)
			t.Error(msg)
		}
	}
}

func TestContinueAfterDuplicateChar(t *testing.T) {
	// cases := map[string]int{
	// 	"anviaj": 5,
	// "anviajbc":        7,
	// "anviajbcapolice": 6,
	// }
	var expected interface{}
	var got interface{}
	cases := []map[string]interface{}{
		{"str": "anviaj", "length": 5, "winningSubstr": "nviaj"},
	}
	for _, testcase := range cases {
		length, substr := lngthOfLongestSubstring(testcase["str"].(string))
		expected = testcase["winningSubstr"]
		got = substr.String()
		if got.(string) != expected.(string) {
			msg := fmt.Sprintf("Expected string to be %s, got %s", expected, got)
			t.Error(msg)
		}

		expected = testcase["length"]
		got = length
		if got.(int) != expected.(int) {
			msg := fmt.Sprintf("Expected length to be %d, got %d", expected, got)
			t.Error(msg)
		}
	}
}
