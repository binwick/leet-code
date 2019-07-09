package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, e := range nums {
		if r, ok := hash[target-e]; ok {
			return []int{r, i}
		}
		hash[e] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	idx := twoSum(nums, 9)
	fmt.Println(idx)
}
