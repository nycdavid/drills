package longestsubstr

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
