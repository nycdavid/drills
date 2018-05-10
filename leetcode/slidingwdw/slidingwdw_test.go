package slidingwdw

import (
	"fmt"
	"testing"
)

func TestSlide(t *testing.T) {
	testCases := []map[string]interface{}{
		{"str": "abcabfg", "res": 5},
		{"str": "anviaj", "res": 5},
		{"str": "abcdefg", "res": 7},
	}
	for _, testCase := range testCases {
		wdw := &SlidingWdw{String: testCase["str"].(string)}
		got := wdw.LengthOfLongest()
		expected := testCase["res"].(int)
		if got != expected {
			msg := fmt.Sprintf("Expected %d, got %d", expected, got)
			t.Error(msg)
		}
	}
}
