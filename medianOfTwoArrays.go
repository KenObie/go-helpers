package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lCombined := append(nums1, nums2...)
	sort.Ints(lCombined)
	n := len(lCombined)

	if n%2 == 0 {
		return (float64(lCombined[n/2] + lCombined[(n/2)-1])) / 2
	} else {
		return float64(lCombined[(n-1)/2])
	}
}

func main() {

}
