package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(threeSum([]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}))
}
//sort自带排序算法超出内存限制。。。
func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for index, v := range nums {
		if v > 0 {
			break
		}
		if (index == 0 || (index > 0 && nums[index] != nums[index-1])) {
			index1, index2 := index+1, len(nums)-1
			for ; index1 < index2; {
				if nums[index1]+nums[index2] == -v {
					result = append(result, []int{v, nums[index1], nums[index2]})
					for ; (index1 < index2 && nums[index1] == nums[index1+1]); {
						index1++
					}
					for ; (index1 < index2 && nums[index2] == nums[index2-1]); {
						index2--
					}
					index1++
					index2--
				} else if nums[index1]+nums[index2] < -v {
					index1++
				} else {
					index2--
				}
			}
		}
	}
	return result
}
