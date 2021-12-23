package main

import (
//	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	dict := make(map[byte]string)
	ss := strings.Split(s, " ")

	for i := range pattern {
		if v, ok := dict[pattern[i]]; ok {
			if v != ss[i] {
				return false
			}
		}else{
			dict[pattern[i]] = ss[i]
		}
	}

	return true
}

/*
func main() {
	pattern, s := []string{"abba", "abba", "aaaa"}, []string{"dog cat cat dog", "dog cat cat fish", "dog cat cat dog"}
	for i := range pattern {
		fmt.Println(wordPattern(pattern[i], s[i]))
	}
}
*/
