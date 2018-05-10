package strcomp

import (
	"bytes"
	"fmt"
)

func MinLen(s1, s2 string) string {
	if len(s1) < len(s2) {
		return s1
	}
	return s2
}

func Compress(str string) string {
	var buf bytes.Buffer
	occs := make(map[rune]int)
	for idx, chr := range str {
		occs[chr]++
		if idx == len(str)-1 {
			buf.WriteString(fmt.Sprintf("%s%d", string(chr), occs[chr]))
			break
		}
		if chr != rune(str[idx+1]) {
			buf.WriteString(fmt.Sprintf("%s%d", string(chr), occs[chr]))
			delete(occs, chr)
		}
	}
	return MinLen(str, buf.String())
}
