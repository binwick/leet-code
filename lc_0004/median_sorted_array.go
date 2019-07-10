package main

import "fmt"

/*
nums1 = [1, 3]
nums2 = [2]
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)

	var nums3 []int
	l := l1 + l2
	odd := true
	if l%2 == 0 {
		odd = false
	}

	var l11, l22 int

	for i := 0; i < l; i++ {
		var i1, i2 int
		if i-l22 < l1 {
			i1 = nums1[i-l22]
		} else {
			nums3 = append(nums3, nums2[i-l11])
			l22++
			continue
		}

		if i-l11 < l2 {
			i2 = nums2[i-l11]
		} else {
			nums3 = append(nums3, nums1[i-l22])
			l11++
			continue
		}

		if i1 < i2 {
			nums3 = append(nums3, i1)
			l11++
		} else {
			nums3 = append(nums3, i2)
			l22++
		}
	}
	fmt.Println(nums3)
	if odd {
		return float64(nums3[l/2])
	} else {
		return float64(nums3[l/2]+nums3[l/2-1]) / 2
	}
}

func main() {
	nums1 := []int{1}
	nums2 := []int{2}
	f := findMedianSortedArrays(nums1, nums2)
	fmt.Println(f)
}
