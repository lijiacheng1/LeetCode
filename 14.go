package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog","sdfsdf"}))
}

func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	for index, v := range strs[0] {
		for i := 0; i < strsLen; i++ {
			if index >= len(strs[i]) || v != rune(strs[i][index]) {
				return strs[0][:index]
			}
		}
		if index == len(strs[0])-1{
			return strs[0]
		}
	}
	return ""
}
