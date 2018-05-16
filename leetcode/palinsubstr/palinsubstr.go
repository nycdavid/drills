package palinsubtr

import (
	"fmt"
)

func longestPalindrome(s string) string {
	s1 := s
	s2 := Reverse(s)
	fmt.Println(s1, s2)
	return ""
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
