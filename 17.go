package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(letterCombinations(""))
}

func letterCombinations(digits string) []string {
	result := []string{""}
	rem := [][]rune{[]rune{'a','b','c'},[]rune{'d','e','f'},[]rune{'g','h','i'},[]rune{'j','k','l'},[]rune{'m','n','o'},[]rune{'p','q','r','s'},[]rune{'t','u','v'},[]rune{'w','x','y','z'}}
	for _,v := range digits{
		if v == '1'{
			continue
		}
		tem := result
		result = []string{}
		for _,x := range tem{
			index,_ := strconv.Atoi(string(v))
			for _,y := range rem[index-2]{
				result = append(result,x+string(y))
			}
		}
	}
	if len(result)<3{
		return []string{}
	}
	return result
}
