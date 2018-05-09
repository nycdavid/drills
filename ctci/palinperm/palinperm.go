package palinperm

import (
	"strings"
)

func PalindromePermutation(str string) bool {
	// generating the different possible permutations of str
	tokens := strings.Split(str, " ")
	joined := strings.Join(tokens, "")
	downcased := strings.ToLower(joined)
	lngth := len(downcased)

	occs := make(map[rune]int)
	for _, chr := range downcased {
		occs[chr]++
	}

	if lngth%2 == 0 {
		for _, v := range occs {
			if v%2 != 0 {
				return false
			}
		}
	} else {
		var foundOdd bool
		for _, v := range occs {
			if v%2 == 1 {
				if foundOdd {
					return false
				} else {
					foundOdd = true
				}
			}
		}
	}

	return true
}
