package main

import (
	"fmt"
)

func main() {
	s := "ab"
	fmt.Println(convert(s, 1))
}

func convert(s string, numRows int) string {
	splites := make(map[int][]byte)
	result := []byte{}
	dir := 1
	row := 0
	for i := 0; i < len(s); i++ {
		splites[row] = append(splites[row], s[i])
		if row == 0 {
			dir = 1
		}
		if row == numRows-1 {
			dir = -1
		}
		if numRows == 1 {
			dir = 0
		}
		row += dir
	}
	for i := 0; i < numRows; i++ {
		result = append(result, splites[i]...)
	}
	return string(result)
}
