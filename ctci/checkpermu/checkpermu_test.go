package checkpermu

import (
	"fmt"
	"testing"
)

func TestIsPermutation(t *testing.T) {
	testCases := []map[string]interface{}{
		{"str1": "abc", "str2": "cba", "res": true},
		{"str1": "abc", "str2": "def", "res": false},
	}

	for _, testCase := range testCases {
		got := IsPermutation(
			testCase["str1"].(string),
			testCase["str2"].(string),
		)
		expected := testCase["res"].(bool)
		if got != expected {
			msg := fmt.Sprintf("Expected %t, got %t", expected, got)
			t.Error(msg)
		}
	}
}

func TestIsPermutationBySort(t *testing.T) {
	testCases := []map[string]interface{}{
		{"str1": "abc", "str2": "cba", "res": true},
		{"str1": "abc", "str2": "def", "res": false},
	}

	for _, testCase := range testCases {
		got := IsPermutationBySort(
			testCase["str1"].(string),
			testCase["str2"].(string),
		)
		expected := testCase["res"].(bool)
		if got != expected {
			msg := fmt.Sprintf("Expected %t, got %t", expected, got)
			t.Error(msg)
		}
	}
}
