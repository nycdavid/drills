package longestsubstr_bg

import (
	"fmt"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	res := lengthOfLongestSubstring("anviaj")

	if res != 5 {
		msg := fmt.Sprintf("Expected length to be %d, got %d", 5, res)
		t.Error(msg)
	}
}
