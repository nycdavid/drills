package palinperm

import (
	"fmt"
	"testing"
)

// 1. Are we counting punctation? Yes
// 2. Are we counting whitespace? No
// 3. Is this case-insensitive?

func TestPalindromePermutation(t *testing.T) {
	testCases := []map[string]interface{}{
		{"str": "Tact Coa", "res": true},
		{"str": "rrrrrr", "res": true},
		{"str": "Raddar", "res": true},
		{"str": "Raddat", "res": false},
		{"str": "aadda", "res": false},
		{"str": "rrrrr", "res": true},
	}

	for _, testCase := range testCases {
		testStr := testCase["str"].(string)
		got := PalindromePermutation(testStr)
		expected := testCase["res"].(bool)

		if got != expected {
			msg := fmt.Sprintf("Expected %s to be %t, got %t", testStr, expected, got)
			t.Error(msg)
		}
	}
}
