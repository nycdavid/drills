package urlify

import (
	"bytes"
	"fmt"
)

func Urlify(str string, length int) string {
	var buf bytes.Buffer
	for i := 0; i < length; i++ {
		chr := str[i]
		if chr == ' ' {
			buf.WriteString("%20")
		} else {
			buf.WriteString(string(chr))
		}
	}
	return buf.String()
}

func UrlifyWithoutBytes(str string, length int) string {
	var chars []rune
	for i := 0; i < length; i++ {
		el := str[i]
		if el == ' ' {
			chars = append(chars, rune('%'))
			chars = append(chars, rune('2'))
			chars = append(chars, rune('0'))
		} else {
			chars = append(chars, rune(el))
		}
	}
	return string(chars)
}
