package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(1994))
}

//一个非常巧妙地方法，免得写出很长的代码
func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}

/*func intToRoman(num int) string {
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
}*/
