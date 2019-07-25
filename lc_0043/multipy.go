package main

import (
	"fmt"
	"math"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	var sum int64
	for i := len(num1) - 1; i >= 0; i-- {
		var next int
		for j := len(num2) - 1; j >= 0; j-- {
			ni, _ := strconv.Atoi(string(num1[i]))
			nj, _ := strconv.Atoi(string(num2[j]))
			tmp := ni*nj + next
			p := math.Pow10(len(num1) - 1 - i + len(num2) - 1 - j)
			tmp = tmp * int(p)
			sum += int64(tmp)
		}
	}
	fmt.Println(sum)
	return ""
}

func main() {

	s := multiply("123", "456")
	fmt.Println(s)
}
