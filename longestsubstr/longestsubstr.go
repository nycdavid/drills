package longestsubstr

import ()

func lngthOfLongestSubstring(str string) (int, *Substring) {
	substr := &Substring{}
	recordHolder := substr
	for _, chr := range str {
		el := string(chr)
		if !substr.Exists(el) {
			substr.Add(el)
		} else {
			if substr.Count() >= recordHolder.Count() {
				recordHolder = substr
			}
			substr = &Substring{}
			idxOfDupe, _ := recordHolder.Find(el)
			startIdx := idxOfDupe + 1
			substr.els = recordHolder.els[startIdx:len(recordHolder.els)]
			substr.Add(el)
		}
	}
	if recordHolder.Count() > substr.Count() {
		return recordHolder.Count(), recordHolder
	} else {
		return substr.Count(), substr
	}
}
