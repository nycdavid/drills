package longestpalinsubtr

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []map[string]string{
		{"input": "babad", "output": "bab"},
		{"input": "cbbd", "output": "bb"},
	}
	for _, c := range cases {
		got := longestPalindrome(c["input"])
		expected := c["output"]

		if got != expected {
			msg := fmt.Sprintf("Expected %s, got %s", expected, got)
			t.Error(msg)
		}
	}
}
