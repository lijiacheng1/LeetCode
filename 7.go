package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	reversedNumber := 0
	for x != 0 {
		reversedNumber = reversedNumber*10 + (x % 10)
		x = x / 10
	}
	if reversedNumber <= -1<<31 || reversedNumber >= 1<<31-1 {
		return 0
	}
	return reversedNumber
}
