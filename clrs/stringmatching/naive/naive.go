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

func (nm *NaiveMatcher) Find(pattern string) (bool, int) {
	t := nm.lenT()
	p := len(pattern)
	for s := 0; s < t-p; s++ {
		if pattern[0:p] == nm.str[s:s+p] {
			nm.p.Println(fmt.Sprintf("Pattern occurs with shift %d", s))
		}
	}
	return false, 0
}

func (nm *NaiveMatcher) lenT() int {
	return len(nm.str)
}
