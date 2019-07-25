package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var s string
	l1 := len(num1)
	l2 := len(num2)
	var a = make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			ni := int(num1[i] - '0') // '0' = 48
			nj := int(num2[j] - '0')
			index := i + j + 1
			a[index] += ni * nj
			if a[index] >= 10 {
				a[index-1] += a[index] / 10 //进位
				a[index] = a[index] % 10
			}
		}
	}
	for i := 0; i < l1+l2; i++ {
		s += strconv.Itoa(a[i])
	}

	if s[0] == '0' {
		return s[1:]
	}

	return s
}

func main() {

	s := multiply("999", "999")
	//s = multiply("498828660196", "840477629533")
	fmt.Println(s)
}
