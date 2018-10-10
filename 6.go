package main

import (
	"fmt"
)

func main() {
	s := "ab"
	fmt.Println(convert(s, 1))
}
//按顺序添加字符到n个子串中,耗时20ms，超过50%的人
func convert(s string, numRows int) string {
	splites := make(map[int][]byte)
	result := []byte{}
	dir := 1
	row := 0
	if numRows == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		splites[row] = append(splites[row], s[i])
		if row == 0 {
			dir = 1
		}else if row == numRows-1 {
			dir = -1
		}
		row += dir
	}
	for i := 0; i < numRows; i++ {
		result = append(result, splites[i]...)
	}
	return string(result)
}



