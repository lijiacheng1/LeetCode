package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

//将每一个字母最后出现的位置保存，每当出现重复字母时，index-index2+1即为子串长度
//输出子串长度最大值
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
