package isunique

import ()

func isUnique(testStr string) bool {
	chars := make([]bool, 128)
	for _, chr := range testStr {
		if chars[chr] == true {
			return false
		}
		chars[chr] = true
	}
	return true
}
