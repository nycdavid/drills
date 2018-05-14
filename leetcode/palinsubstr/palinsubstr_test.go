package palinsubtr

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []map[string]string{
		{"str": "b", "res": "b"},
		{"str": "babad", "res": "bab"},
		{"str": "cbbd", "res": "bb"},
	}

	for _, c := range cases {
		got := longestPalindrome(c["str"])
		expected := c["res"]

		if got != expected {
			msg := fmt.Sprintf("Expected %s, got %s", expected, got)
			t.Error(msg)
		}
	}
}
