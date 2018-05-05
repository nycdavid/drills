package longestsubstr

import ()

func lngthOfLongestSubstring(str string) (int, *Substring) {
	substr := &Substring{}
	for _, chr := range str {
		el := string(chr)
		if !substr.Exists(el) {
			substr.Add(el)
		}
	}
	// newSub := &Substring{}
	// recordHolder := newSub
	//
	// for _, chr := range str {
	// 	if !newSub.Exists(string(chr)) {
	// 		newSub.Add(string(chr))
	// 	} else {
	// 		delete(newSub.els, string(chr))
	// 		newSub.Add(string(chr))
	// 	}
	// 	if newSub.Count() >= recordHolder.Count() {
	// 		recordHolder = newSub
	// 	}
	// }
	// return recordHolder.Count()
	return substr.Count(), substr
}
