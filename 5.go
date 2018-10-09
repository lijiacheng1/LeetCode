package main

import (
	"fmt"
)

func main() {
	s := "ac"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	splits := []byte(s)
	max := 0
	pre, last := 0, 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			length := 2
			pre1 := i
			last1 := i + 1
			for ; pre1 >= 0 && last1 < len(s); {
				if s[pre1] == s[last1] {
					pre1--
					last1++
					length += 2
				} else {
					break
				}
			}
			if length > max {
				max = length
				pre = pre1 + 1
				last = last1
			}
		}
		if i+2 < len(s) && s[i] == s[i+2] {
			length := 3
			pre1 := i
			last1 := i + 2
			for ; pre1 >= 0 && last1 < len(s); {
				if s[pre1] == s[last1] {
					pre1--
					last1++
					length += 2
				} else {
					break
				}
			}
			if length > max {
				max = length
				pre = pre1 + 1
				last = last1
			}
		}
	}
	if max == 0 && len(s) >= 1 {
		max = 1
		pre = 0
		last = 1
	}
	return string(splits[pre:last])
}
