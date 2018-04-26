package twosum

import ()

func TwoSum(collection []int, target int) []int {
	hsh := make(map[int]int)
	for idx, num := range collection {
		hsh[num] = idx
	}

	for idx, num := range collection {
		complement := target - num
		if _, ok := hsh[complement]; ok {
			if hsh[complement] != idx {
				pair := []int{hsh[complement], hsh[num]}
				return pair
			}
		}
	}
	return []int{0, 0}
}
