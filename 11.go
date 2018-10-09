package main

import (
	"fmt"
)

func main() {
	asd := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(asd))
}

func maxArea(height []int) int {
	i :=0
	j:= len(height)-1
	maxWater := 0
	for ;i!=j ;  {
		maxWater = Max(Min(height[i],height[j]) * (j-i),maxWater)
		if height[i]<height[j] {
			i++
		}else {
			j--
		}
	}
	return maxWater
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}