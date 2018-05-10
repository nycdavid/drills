package strcomp

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []map[string]string{
		{"str": "aabcccccaaa", "res": "a2b1c5a3"},
		{"str": "abc", "res": "abc"},
	}

	for _, test := range tests {
		got := Compress(test["str"])
		expected := test["res"]

		if got != expected {
			msg := fmt.Sprintf("Expected %s, got %s", expected, got)
			t.Error(msg)
		}
	}
}
