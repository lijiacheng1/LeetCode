package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2,-1,-1,-1}, 0))
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for index, val := range nums {
		if (index == 0 || index > 0 && nums[index] != nums[index-1]) {
			for index0 := index + 1; index0 < len(nums)-2; index0++ {
				if (index0 == index+1 ||nums[index0] != nums[index0-1]) {
					index1, index2 := index0+1, len(nums)-1
					for ; index1 < index2; {
						rem := target - val - nums[index0]
						if nums[index1]+nums[index2] == rem {
							result = append(result, []int{val, nums[index0], nums[index1], nums[index2]})
							for ; (index1 < index2 && nums[index1] == nums[index1+1]); {
								index1++
							}
							for ; (index1 < index2 && nums[index2] == nums[index2-1]); {
								index2--
							}
							index1++
							index2--
						} else if nums[index1]+nums[index2] < rem {
							index1++
						} else {
							index2--
						}
					}
				}
			}
		}
	}
	return result
}
