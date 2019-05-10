package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))

}

func removeOuterParentheses(S string) string {
	count := 0
	tstring := make([]rune, len(S))
	i := 0
	for _, v := range S {
		if v == '(' {
			count++
			if count == 1 {
			} else {
				tstring[i] = v
				i++
			}
		} else { // v == ')'
			count--
			if count == 0 {
			} else {
				tstring[i] = v
				i++
			}
		}
	}

	return string(tstring[0:i])
}
