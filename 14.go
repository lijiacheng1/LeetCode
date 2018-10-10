package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog","sdfsdf"}))
}
//对于第一个字符串的每个字母，与其他字符串的相应字母进行比较，这么笨的方法居然也能100%
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
