package sortedmedian

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	testCases := []map[string]interface{}{
		{"nums1": []int{1, 3}, "nums2": []int{2}, "sol": 2.0},
		{"nums1": []int{1, 2}, "nums2": []int{3, 4}, "sol": 2.5},
	}

	for _, testCase := range testCases {
		nums1 := testCase["nums1"].([]int)
		nums2 := testCase["nums2"].([]int)
		got := findMedianSortedArrays(nums1, nums2)
		expected := testCase["sol"].(float64)

		if got != expected {
			msg := fmt.Sprintf("Expected %g, got %g", expected, got)
			t.Error(msg)
		}
	}
}
