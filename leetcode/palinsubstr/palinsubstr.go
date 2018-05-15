package palinsubtr

import (
	"fmt"
)

func longestPalindrome(s string) string {
	caps := make([]int, len(s))
	for idx := range s {
		i := idx
		j := idx
		ln := 1
		if i == 0 || j == len(s)-1 {
			caps[idx] = ln
			continue
		}
		i--
		j++
		if s[i] == s[j] {
			ln = ln + 2
			caps[idx] = ln
		}
	}
	return ""
}
