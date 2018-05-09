package longestsubstr

import ()

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lngthOfLongestSubstring(str string) (int, *Substring) {
	lngth := len(str)
	occs := make(map[rune]bool)
	var i, j, ans int
	for i < lngth && j < lngth {
		chr := str[j]
		if occs[rune(chr)] == false {
			occs[rune(chr)] = true
			ans = Max(ans, j-i+1)
			j++
		} else {
			i++
		}
	}
	ans = Max(ans, len(occs))
	return ans, &Substring{}
}

func allUnique(str string) bool {
	hsh := make(map[rune]int)
	for _, chr := range str {
		hsh[chr]++
	}
	for _, ct := range hsh {
		if ct != 1 {
			return false
		}
	}
	return true
}
