package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{[][]}"))
}

func isValid(s string) bool {
	rem := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			rem = append(rem, v)
		} else {
			if v == ')' {
				if len(rem) != 0 && rem[len(rem)-1] == '(' {
					rem = rem[:len(rem)-1]
				} else {
					return false
				}
			}
			if v == ']' {
				if len(rem) != 0 && rem[len(rem)-1] == '[' {
					rem = rem[:len(rem)-1]
				} else {
					return false
				}
			}
			if v == '}' {
				if len(rem) != 0 && rem[len(rem)-1] == '{' {
					rem = rem[:len(rem)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(rem)==0{
		return true
	}
	return false
}
