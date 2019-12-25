package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	d1 := make(map[string]int)
	d2 := make(map[string]int)

	for _, e := range s {
		d1[string(e)] = d1[string(e)] + 1
	}
	for _, e := range t {
		d2[string(e)] = d2[string(e)] + 1
	}
	return reflect.DeepEqual(d1, d2)
}

func main() {
	anagram := isAnagram("tar", "rat")
	fmt.Println(anagram)
}
