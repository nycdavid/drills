package longestsubstr_bg

import (
	"fmt"
)

func lengthOfLongestSubstring(substr string) int {
	n := len(substr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			testSubstr := substr[i:j]
			fmt.Println(testSubstr)
			if allUnique(testSubstr) {
				return len(testSubstr)
			}
		}
	}
	return 0
}

func allUnique(str string) bool {
	var chars map[string]int
	for _, chr := range str {
		el := string(chr)
		if chars[el] != 0 {
			chars[el] = 1
		} else {
			return false
		}
	}
	return true
}
