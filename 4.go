package main

import (
	"fmt"
	"math"
)

func main() {
	var s = []int{2, 7, 15, 10, 1};
	fmt.Println(twoSum(s, 8));
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid := len(nums1) + len(nums2)
	if mid%2 == 0 {

	} else {
		medain := math.Floor(float64(mid/2)) + 1

	}
}
