package urlify

import (
	"fmt"
	"testing"
)

func TestUrlify(t *testing.T) {
	testStr := "Mr John Smith   "
	got := Urlify(testStr, 13)
	expected := "Mr%20John%20Smith"

	if got != expected {
		msg := fmt.Sprintf("Expected %s, got %s", expected, got)
		t.Error(msg)
	}
}

func TestUrlifyWithoutBytes(t *testing.T) {
	testStr := "Mr John Smith   "
	got := UrlifyWithoutBytes(testStr, 13)
	expected := "Mr%20John%20Smith"

	if got != expected {
		msg := fmt.Sprintf("Expected \"%s\", got \"%s\"", expected, got)
		t.Error(msg)
	}
}
