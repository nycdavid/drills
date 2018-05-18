package naive

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type BufferPrinter struct {
	buf *bytes.Buffer
}

func (bp BufferPrinter) Println(a ...interface{}) (int, error) {
	for _, s := range a {
		bp.buf.WriteString(s.(string))
	}
	return fmt.Println(a)
}

func TestNaiveStringMatcher(t *testing.T) {
	var buf bytes.Buffer
	bp := BufferPrinter{buf: &buf}
	mat := &NaiveMatcher{str: "abcdefgbc", p: bp}
	exists := mat.Find("bc")

	if exists != true {
		t.Error("Expected to find substring")
	}
	expected := "Pattern occurs with shift 1\nPattern occurs with shift 7"
	got := buf.String()
	if strings.TrimSpace(got) != strings.TrimSpace(expected) {
		msg := fmt.Sprintf("Expected string to be %s, got %s", expected, got)
		t.Error(msg)
	}
}

func Test32_1_1(t *testing.T) {
	T := "000010001010001"
	p := "0001"
	var buf bytes.Buffer
	bp := BufferPrinter{buf: &buf}
	mat := &NaiveMatcher{str: T, p: bp}

	exists := mat.Find(p)
	expected := "Pattern occurs with shift 1\nPattern occurs with shift 5\nPattern occurs with shift 11"
	got := buf.String()

	if exists != true {
		t.Error("Expected to find substring")
	}
	if strings.TrimSpace(got) != strings.TrimSpace(expected) {
		msg := fmt.Sprintf("Expected string to be %s, got %s", expected, got)
		t.Error(msg)
	}
}
