package checkpermu

import (
	"sort"
)

type SortableRunes []rune

func (sr SortableRunes) Len() int {
	return len(sr)
}

func (sr SortableRunes) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func (sr SortableRunes) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func IsPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	// all strings of the same size
	chars := make([]int, 128)
	for _, chr := range str1 {
		chars[chr] = 1
	}
	for _, chr := range str2 {
		if chars[chr] != 1 {
			return false
		}
	}
	return true
}

func IsPermutationBySort(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	s1Runes := []rune(str1)
	s2Runes := []rune(str2)
	sort.Sort(SortableRunes(s1Runes))
	sort.Sort(SortableRunes(s2Runes))
	for idx, rne := range s1Runes {
		if rne != s2Runes[idx] {
			return false
		}
	}
	return true
}
