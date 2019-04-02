package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(5))
}

func generateParenthesis(n int) []string {
	result := &[]string{}
	backtrack22(result, "", 0, 0, n)
	return *result
}

func backtrack22(result *[]string, str string, open int, close int, max int) {
	if len(str) >= max<<1 {
		*result = append(*result, str)
		return
	}
	if open < max {
		backtrack22(result, str+"(", open+1, close, max)
	}
	if close < open {
		backtrack22(result, str+")", open, close+1, max)
	}
}
