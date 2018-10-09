package main

import (
	"fmt"
)

func main() {
	var s = []int{2, 7, 15, 10, 1}

	fmt.Println(twoSum(s, 8))
}

func twoSum(nums []int, target int) []int {
	numsmap := make(map[int]int)
	for index, v := range nums {
		if _, ok := numsmap[target-v]; ok {
			return []int{numsmap[target-v], index}
		}
		numsmap[v] = index
	}
	return []int{};
}
