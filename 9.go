package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(1796776971))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	reversedNumber := 0
	for y != 0 {
		reversedNumber = reversedNumber*10 + y%10
		y = y / 10
	}

	if reversedNumber == x {
		return true
	}

	return false
}