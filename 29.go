package main

import "fmt"

func main() {
	fmt.Println(divide(7, -3))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func divide(dividend int, divisor int) int {
	endNev, orNev := false, false
	result :=0
	if dividend <= 0 {
		dividend = -dividend
		endNev = true
	}
	if divisor <= 0 {
		divisor = -divisor
		orNev = true
	}
	for i := dividend; i >= divisor; i -= divisor {
		result++
	}
	if endNev != orNev{
		result = -result
	}
	return result
}
