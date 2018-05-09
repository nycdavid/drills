package oneaway

import (
	"fmt"
	"testing"
)

func TestOneAway(t *testing.T) {
	testCases := []map[string]interface{}{
		{"start": "pale", "fin": "pale", "res": true},
		{"start": "pale", "fin": "bale", "res": true},
		{"start": "pale", "fin": "ple", "res": true},
		{"start": "pale", "fin": "pales", "res": true},
	}

	for _, testCase := range testCases {
		got := OneAway(testCase["start"].(string), testCase["fin"].(string))
		expected := testCase["res"].(bool)
		if expected != got {
			msg := fmt.Sprintf(
        "Expected %s -> %s to be %t, got %t",
        testCase["start"].(string),
        testCase["fin"].(string),
        expected,
        got,
      )
			t.Error(msg)
		}
	}
}
