package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var subStr string
	var maxLength int

	for _, e := range s {
		index := strings.Index(subStr, string(e))
		if index == -1 {

		} else if index == 0 {
			subStr = subStr[1:]
		} else if index == len(subStr)-1 {
			subStr = ""
		} else {
			subStr = subStr[index+1:]
		}
		subStr += string(e)
		fmt.Println(subStr)
		if len(subStr) > maxLength {
			maxLength = len(subStr)
		}
	}
	return maxLength
}

func main() {
	var s string
	s = "abcabcbb"
	//s = "bbbbb"
	//s = "aab"
	s = "aabaab!bb"
	//s = "dvdf"
	//s = "pwwkew"
	//s = "ohvhjdml"
	l := lengthOfLongestSubstring(s)

	fmt.Println(l)
}
