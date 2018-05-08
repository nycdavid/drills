package palinperm

import (
	"fmt"
	"sort"
	"strings"
)

type SortableRunes []rune

func (sr SortableRunes) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func (sr SortableRunes) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func (sr SortableRunes) Len() int {
	return len(sr)
}

func PalindromePermutation(str string) bool {
	// generating the different possible permutations of str
	tokens := strings.Split(str, " ")
	joined := strings.Join(tokens, "")
	downcased := strings.ToLower(joined)
	lngth := len(downcased)

	occs := make([]rune, lngth)
	for idx, chr := range downcased {
		occs[idx] = chr
	}
  for idx, code := range occs {
    code.
  }

	if lngth%2 == 0 {
	} else {
	}

	fmt.Println("")
	return true
}
