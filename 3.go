package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	stringmap := make(map[rune]int)
	index2 := 0
	for index, v := range s {
		if _, ok := stringmap[v]; ok {
			if index2 < stringmap[v]+1 {
				index2 = stringmap[v] + 1
			}
		}
		stringmap[v] = index
		if index-index2+1 > max {
			max = index - index2 + 1
		}
	}
	return max
}
