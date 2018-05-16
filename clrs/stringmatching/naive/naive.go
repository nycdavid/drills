package naive

type Printer interface {
	Println(a ...interface{}) (int, error)
}

type NaiveMatcher struct {
	str string
	p   Printer
}

func (nm *NaiveMatcher) Find(substr string) (bool, int) {
	return false, 0
}
