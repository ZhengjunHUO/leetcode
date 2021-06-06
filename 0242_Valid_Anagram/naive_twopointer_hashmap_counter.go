package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false 
	}

	m := make(map[rune]int) 
	for _,v := range s {
		m[v] += 1
	}
	count := len(s)

	rp := 0
	tr := []rune(t)

	for rp < len(t) {
		if m[tr[rp]] > 0 {
			count -= 1
		}
		m[tr[rp]] -= 1 
		rp += 1

		if count == 0 {
			return true
		}
	}
	return false
}

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))

	s, t = "rat", "car"
	fmt.Println(isAnagram(s, t))
}
