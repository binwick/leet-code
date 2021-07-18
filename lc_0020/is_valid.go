package main

import "fmt"

var parenMap = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func isValid(s string) bool {
	var stack []string
	for _, e := range s {
		c := string(e)
		if _, ok := parenMap[c]; !ok {
			// 左括号
			stack = append(stack, c)
		} else {
			// 右括号
			if len(stack) == 0 {
				return false
			}
			// pop
			pop := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if parenMap[c] != pop {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	valid := isValid("()")
	fmt.Println(valid)
	valid = isValid("(})")
	fmt.Println(valid)
	valid = isValid("({})")
	fmt.Println(valid)
	valid = isValid("()[]{}")
	fmt.Println(valid)
}
