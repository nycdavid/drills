package palinsubtr

import (
	"strings"
)

func longestPalindrome(s string) string {
	var j int
	var ans, substr string
	for j < len(s) {
		substr = s[0 : j+1]
		if isPalindrome(substr) {
			ans = MaxLen(ans, substr)
		}
		j++
	}
	for i := range s {
		substr = s[i:len(s)]
		if isPalindrome(substr) {
			ans = MaxLen(ans, substr)
		}
	}
	return ans
}

func isPalindrome(s string) bool {
	lgth := len(s)
	if lgth == 1 {
		return true
	}
	var rev []string
	for idx := range s {
		revIdx := lgth - 1 - idx
		rev = append(rev, string(s[revIdx]))
	}
	if strings.Join(rev, "") == s {
		return true
	}
	return false
}

func MaxLen(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
