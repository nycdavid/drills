package twosum

import (
	"testing"
)

// function that accepts a slice and a target
// returns the indices of the array such that the ints add up to the target
func TestTwoSum(t *testing.T) {
	testCases := []map[string]interface{}{
		{"testSlc": []int{1, 2, 3}, "target": 5, "idxs": []int{1, 2}},
		{"testSlc": []int{3, 2, 5}, "target": 5, "idxs": []int{0, 1}},
		{"testSlc": []int{4, 1, 10}, "target": 5, "idxs": []int{0, 1}},
	}

	for _, testCase := range testCases {
		testSlc := testCase["testSlc"].([]int)
		target := testCase["target"].(int)
		received := TwoSum(testSlc, target)

		solns := testCase["idxs"].([]int)
		if solns[0] != received[0] && solns[0] != received[1] {
			t.Error("Incorrect result")
		}
		if solns[1] != received[0] && solns[1] != received[1] {
			t.Error("Incorrect result")
		}
	}
}
