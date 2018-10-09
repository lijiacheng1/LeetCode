package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	pre,last,result := 0,0,0
	for _, v := range s {
		pre = last
		last = findV(v)
		if pre<last{
			result -= 2*pre
		}
		result += last
	}
	return result
}
func findV(v rune) int {
	switch v {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
