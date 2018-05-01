package longestsubstr

import ()

type Substring struct {
	els map[string]int
}

func (sub *Substring) Exists(el string) bool {
	if sub.els[el] == 0 {
		return false
	}
	return true
}

func (sub *Substring) Add(el string) error {
	if sub.els == nil {
		sub.els = make(map[string]int)
	}
	sub.els[el] = 1
	return nil
}

func (sub *Substring) Count() int {
	return len(sub.els)
}

func lengthOfLongestSubstring(str string) int {
	recordHolder := &Substring{}
	newSub := &Substring{}
	for _, chr := range str {
		if !newSub.Exists(string(chr)) {
			newSub.Add(string(chr))
		} else {
			if newSub.Count() >= recordHolder.Count() {
				recordHolder = newSub
				newSub = &Substring{}
			}
		}
	}
	return recordHolder.Count()
}