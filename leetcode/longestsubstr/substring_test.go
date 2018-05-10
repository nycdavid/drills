package longestsubstr

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	substr := &Substring{}

	if substr.Exists("a") != false {
		t.Error("Expected false")
	}
	substr.Add("a")
	substr.Add("b")
	if substr.Exists("b") != true {
		t.Error("Expected true")
	}
	if substr.Exists("a") != true {
		t.Error("Expected true")
	}
}

func TestAdd(t *testing.T) {
	substr := &Substring{}
	substr.Add("f")

	if len(substr.els) != 1 {
		msg := fmt.Sprintf("Expected %d elements, got %d", 1, len(substr.els))
		t.Error(msg)
	}
}

func TestLast(t *testing.T) {
	substr := &Substring{}
	substr.Add("f")
	substr.Add("o")
	substr.Add("o")

	expected := "o"
	got := substr.Last()

	if got != expected {
		msg := fmt.Sprintf("Expected %s, got %s", expected, got)
		t.Error(msg)
	}
}

func TestLastWhenEmpty(t *testing.T) {
	substr := &Substring{}

	if substr.Last() != "" {
		t.Error("Expected \"\"")
	}
}

func TestCount(t *testing.T) {
	substr := &Substring{}
	substr.Add("a")
	expected := 1
	got := substr.Count()

	if got != expected {
		msg := fmt.Sprintf("Expected %d, got %d", expected, got)
		t.Error(msg)
	}
}

func TestString(t *testing.T) {
	substr := &Substring{}
	substr.Add("a")
	substr.Add("b")
	substr.Add("c")

	expected := "abc"
	got := substr.String()
	if got != expected {
		msg := fmt.Sprintf("Expected %s, got %s", expected, got)
		t.Error(msg)
	}
}

func TestFind(t *testing.T) {
	substr := &Substring{}
	substr.Add("a")
	substr.Add("b")

	expected := 1
	got, err := substr.Find("b")
	if err != nil {
		t.Error(err)
	}

	if expected != got {
		msg := fmt.Sprintf("Expected index to be %d, got %d", expected, got)
		t.Error(msg)
	}
}
