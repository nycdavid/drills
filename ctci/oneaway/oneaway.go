package oneaway

// OneAway is a function that accepts two strings and returns true
// if they are (at most) one edit away
func OneAway(start, fin string) bool {
	startLen := len(start)
	finLen := len(fin)

	if startLen == finLen {
		// test for single replacement
		return isOneOrLess(start, fin)
	} else if startLen-1 == finLen {
		return isSingleRmvl(start, fin)
		// test for single removal
	} else if startLen+1 == finLen {
		// test for single insertion
	}
	return false
}

func isOneOrLess(start, fin string) bool {
	charDiffs := 0
	if start == fin {
		return true
	}
	for idx, chr := range start {
		if chr != rune(fin[idx]) {
			charDiffs++
		}
	}
	if charDiffs == 1 {
		return true
	}
	return false
}

func isSingleRmvl(start, fin string) bool {
	return false
}
