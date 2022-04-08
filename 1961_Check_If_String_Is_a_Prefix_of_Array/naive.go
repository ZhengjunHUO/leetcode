package main

import (
	"fmt"
)

func isPrefixString(s string, words []string) bool {
	var idx int
	for i:=0; i<len(words) && idx<len(s); i++ {
		for j := range words[i] {
			// s reach the end but words[i] not => not a prefix, or
			// the letter don't match
			if idx == len(s) || words[i][j] != s[idx] {
				return false
			}
			idx++
		}
	}

	// if all words is compared with s but s still has letters left => not a prefix
	return idx == len(s)
}

func main() {
	s := "iloveleetcode"
	words := [][]string{[]string{"i","love","leetcode","apples"}, []string{"apples","i","love","leetcode"}}

	for i := range words {
		fmt.Println(isPrefixString(s, words[i]))
	}
}
