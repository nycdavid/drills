package longestsubstr

import (
	"errors"
	"strings"
)

type Substring struct {
	els []string
}

func (sub *Substring) Exists(testEl string) bool {
	for _, el := range sub.els {
		if el == testEl {
			return true
		}
	}
	return false
}

func (sub *Substring) Add(testEl string) {
	sub.els = append(sub.els, testEl)
}

func (sub *Substring) Last() string {
	length := len(sub.els)
	if length == 0 {
		return ""
	} else {
		return sub.els[length-1]
	}
}

func (sub *Substring) Count() int {
	return len(sub.els)
}

func (sub *Substring) String() string {
	return strings.Join(sub.els, "")
}

func (sub *Substring) Find(el string) (int, error) {
	if len(sub.els) == 0 {
		return -1, errors.New("sub.els is empty")
	}
	for idx, char := range sub.els {
		if char == el {
			return idx, nil
		}
	}
	return -1, nil
}
