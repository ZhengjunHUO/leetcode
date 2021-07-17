package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	dict := make(map[byte]int) 

	for i := range s {
		dict[s[i]] += 1
	}

	for i := range s {
		if v := dict[s[i]]; v == 1 {
			return i	
		}
	}

	return -1
}

func main() {
	strs := []string{"leetcode", "loveleetcode", "aabb"}

	for _, v := range strs {
		fmt.Println(firstUniqChar(v))
	}
}
