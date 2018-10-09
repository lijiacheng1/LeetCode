package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(myAtoi(" .1"))
}

func myAtoi(str string) int {
	f := func(c rune) bool {
		if c == ' '||c == '.'{
			return true
		}
		return false
	}
	splites := strings.FieldsFunc(str,f)

	for _, i := range splites {
		fmt.Println(i)
		data, _ := strconv.Atoi(string(i))
		if i == "" {
			continue
		} else if data == 0 {
			break
		} else if data != 0 {
			if data <= -1<<31 {
				return -1 <<31
			}else if data >= 1<<31 -1 {
				return 1<<31 - 1
			}else {
				return data
			}
		}
	}
	return 0
}
