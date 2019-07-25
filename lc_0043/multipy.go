package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	var s string
	for i := len(num1) - 1; i >= 0; i-- {
		var next int
		for j := len(num2) - 1; j >= 0; j-- {
			ni := int(num1[i] - 48)
			nj := int(num2[j] - 48)
			tmp := ni*nj + next
			if i == len(num1)-1 {
				s = strconv.Itoa(tmp%10) + s
				next = tmp / 10
			} else {
				sum := int(s[j]-48) + tmp%10
				next = sum/10 + tmp/10
				s = s[0:j] + strconv.Itoa(sum%10) + s[j+1:]
			}
		}
		s = strconv.Itoa(next) + s
	}

	if s[0] == 48 {
		return s[1:]
	}
	return s
}

func main() {

	s := multiply("123", "456")
	s = multiply("498828660196", "840477629533")
	fmt.Println(s)
}
