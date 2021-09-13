package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	var rslt byte

	for i := range s {
		rslt ^= s[i]
	}

	for i := range t {
		rslt ^= t[i]
	}

	return rslt
}

func main() {
	s, t := []string{"abcd", "", "a", "ae"}, []string{"abcde", "y", "aa", "aea"}
	for i := range s {
		fmt.Println(string(findTheDifference(s[i], t[i])))
	}
}
