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
	return 0, nil
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
