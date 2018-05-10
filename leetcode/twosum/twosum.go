package twosum

import ()

func TwoSum(collection []int, target int) []int {
	hsh := make(map[int]int)
	for idx, num := range collection {
		hsh[num] = idx
	}

	for idx, num := range collection {
		// this is the other number we're looking for
		complement := target - num

		// If the value is found in hsh
		if _, ok := hsh[complement]; ok {
			// and it's index is NOT the current index
			if hsh[complement] != idx {
				pair := []int{idx, hsh[complement]}
				return pair
			}
		}
	}
	return []int{0, 0}
}
