package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	result := []string{}
	x := num / 1000
	for i := 0; i < x; i++ {
		result = append(result, "M")
	}
	num = num % 1000
	x = num/100
	yuToRoman(x,"M","D","C",&result)
	num = num%100
	x = num/10
	yuToRoman(x,"C","L","X",&result)
	num = num%10
	yuToRoman(num,"X","V","I",&result)
	return strings.Join(result,"")
}

func yuToRoman(x int, shu10, shu5, shu1 string, result *[]string) {
	switch x {
	case 1:
		*result = append(*result, shu1)
	case 2:
		*result = append(*result, shu1, shu1)
	case 3:
		*result = append(*result, shu1, shu1, shu1)
	case 4:
		*result = append(*result, shu1, shu5)
	case 5:
		*result = append(*result, shu5)
	case 6:
		*result = append(*result, shu5, shu1)
	case 7:
		*result = append(*result, shu5, shu1,shu1)
	case 8:
		*result = append(*result, shu5, shu1, shu1, shu1)
	case 9:
		*result = append(*result,shu1,shu10)
	}
}
