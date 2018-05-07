package isunique

import (
	"fmt"
	"testing"
)

// Implement an algorithm to determine if a string has all unique chars.

func TestIsUnique(t *testing.T) {
	testCases := []map[string]interface{}{
		{"str": "abcde", "res": true},
		{"str": "abca", "res": false},
	}

	for _, testCase := range testCases {
		expected := testCase["res"].(bool)

		got := isUnique(testCase["str"].(string))

		if expected != got {
			msg := fmt.Sprintf("Expected %t, got %t", expected, got)
			t.Error(msg)
		}
	}
}
