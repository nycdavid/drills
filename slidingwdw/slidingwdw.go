package slidingwdw

import ()

type SlidingWdw struct {
	String string
}

func (sw *SlidingWdw) LengthOfLongest() int {
	lngth := len(sw.String)
	var i, j, ans int
	occs := make(map[rune]bool)

	for i < lngth && j < lngth {
		iRne := rune(sw.String[i])
		jRne := rune(sw.String[j])
		if !occs[jRne] {
			occs[jRne] = true
			j++
		} else {
			// We've found a duplication so keep bringing up
			// the lower bound until it's not present anymore
			ans = Max(ans, j-i)
			occs[iRne] = false
			i++
		}
		ans = Max(ans, j-i)
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
