package main

import "fmt"

func main() {
	s := "cbbd"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	var start int
	var end int
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len3 := findMax(len1, len2)
		if len3 > end-start {
			start = i - (len3-1)/2
			end = i + len3/2
		}
	}
	return s[start : end+1]
}

func findMax(len1, len2 int) int {
	if len1 > len2 {
		return len1
	}
	return len2
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
