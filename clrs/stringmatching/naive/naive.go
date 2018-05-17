package naive

import (
	"fmt"
)

type Printer interface {
	Println(a ...interface{}) (int, error)
}

type StdoutPrinter struct{}

func (s StdoutPrinter) Println(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

type NaiveMatcher struct {
	str string
	p   Printer
}

func NewNaiveMatcher(str string, p Printer) *NaiveMatcher {
	return &NaiveMatcher{str: str, p: StdoutPrinter{}}
}

func (nm *NaiveMatcher) Find(pattern string) bool {
	t := nm.lenT()
	p := len(pattern)
	var found bool
	for s := 0; s < t-p+1; s++ {
		if pattern[0:p] == nm.str[s:s+p] {
			nm.p.Println(fmt.Sprintf("Pattern occurs with shift %d\n", s))
			found = true
		}
	}
	return found
}

func (nm *NaiveMatcher) lenT() int {
	return len(nm.str)
}
