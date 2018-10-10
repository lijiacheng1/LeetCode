package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(threeSumClosest([]int{25, -12, 2, 66, -8, 67, 67, -93, -97, -68, -49, -24, 49, 90, 31, 87, -1, 14, 50, -25, 69, -91, -23, 82, -43, 96, 80, 43, 4, -87, 40, 24, -71, -10, -16, 78, -60, -20, -84, 69, 84, 84, 16, -23, -25, 88, 15, 72, -82, -72, -16, 49, 50, 26, 3, 26, -92, 82, -69, 60, -81, 24, -25, -47, -77, -37, -33, 69, 21, -50, 56, -43, 18, -87, 96, 47, 59, 37, 100, 23, -34, -69, -12, -83, -65, 91, 75, 100, 24, 80, 64, -27, -31, 39, -46, -73, 88, -13, -10, 67, 95, 27, 91, -61, -44, 67, 0, -29, -57, -61, -62, 83, -6, 82, -58, -58, -5, -87, -44, 9, -20, -28, 3, 17, 57, -63, 78, 34, 7, -68, 30, -49, 77, -97, 15, -42, -49, -22, -60, 78, 84, 35, 19}, 250))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	distance := 9999999
	result := 0
	for index, v := range nums {
		index1, index2 := index+1, len(nums)-1
		for ; index1 < index2; {
			sum := v + nums[index1] + nums[index2]
			rem := target - sum
			if rem < 0 {
				rem = -rem
			}
			if rem < distance {
				result = sum
				distance = rem
			}
			if sum < target {
				index1++
			} else {
				index2--
			}
		}
	}
	return result
}
