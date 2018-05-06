package sortedmedian

import (
	"math"
	"sort"
)

type SortableInts []int

func (si SortableInts) Len() int {
	return len(si)
}

func (si SortableInts) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}

func (si SortableInts) Less(i, j int) bool {
	return si[i] < si[j]
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var combinedArr []int
	for _, num := range nums1 {
		combinedArr = append(combinedArr, num)
	}
	for _, num := range nums2 {
		combinedArr = append(combinedArr, num)
	}
	sort.Sort(SortableInts(combinedArr))
	median := findMedian(combinedArr)
	// sort array
	return median
}

func findMedian(nums []int) float64 {
	elsCt := len(nums)
	if elsCt%2 == 0 {
		// even number of elements
		quot := float64(elsCt) / 2.0
		idx := int(quot) - 1
		avg := float64((nums[idx] + nums[idx+1])) / 2.0
		return avg
	} else {
		// odd number of elements
		quot := float64(elsCt) / 2.0
		idx := int(math.Ceil(quot) - 1)
		return float64(nums[idx])
	}
	return 0.0
}
