package naive

import (
	"testing"
)

func TestNaiveStringMatcher(t *testing.T) {
	m := &NaiveMatcher{str: "abcdefg"}
	exists, shift := m.Find("bc")

	if exists != true {
		t.Error("Expected to find substring")
	}
	if shift != 1 {
		t.Error("Expected shift to be 1")
	}
}
