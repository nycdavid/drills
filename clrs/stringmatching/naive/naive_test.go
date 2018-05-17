package naive

import (
	"bytes"
	"fmt"
	"testing"
)

type BufferPrinter struct {
	buf *bytes.Buffer
}

func (bp BufferPrinter) Println(a ...interface{}) (int, error) {
	for _, s := range a {
		bp.buf.WriteString(s.(string))
	}
	return 0, nil
}

func TestNaiveStringMatcher(t *testing.T) {
	var buf bytes.Buffer
	bp := BufferPrinter{buf: &buf}
	mat := &NaiveMatcher{str: "abcdefg", p: bp}
	exists, shift := mat.Find("bc")

	if exists != true {
		t.Error("Expected to find substring")
	}
	if shift != 1 {
		t.Error("Expected shift to be 1")
	}
	expected := "Pattern occurs with shift 1"
	got := buf.String()
	if got != expected {
		msg := fmt.Sprintf("Expected string to be %s, got %s", expected, got)
		t.Error(msg)
	}
}
