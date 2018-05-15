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
		for i > 0 && j < len(s) {
			i--
			j++
			if s[i] == s[j] {
				ln += 2
			}
			caps[idx] = ln
		}
		caps[idx] = ln
	}
	fmt.Println(caps)
	return ""
}
