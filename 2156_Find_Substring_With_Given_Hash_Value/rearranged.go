package main

import (
	"fmt"
)

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	var i int
	curr, m := int(s[0]-96), 1
	for i=1; i<len(s); i++ {
		if i < k {
			m *= power
			curr += int(s[i] - 96) * m
			continue
		}

		if i == k && curr%modulo == hashValue {
			return s[i-k:i]
		}

		curr -= int(s[i-k] - 96)
		curr /= power
		curr += int(s[i] - 96) * m

		if curr%modulo == hashValue {
			return s[i-k+1:i+1]
		}
	}

	return ""
}

func main() {
	s := []string{"leetcode", "fbxzaad"}
	power := []int{7,31}
	modulo := []int{20,100}
	k := []int{2,3}
	hashValue := []int{0,32}

	for i := range s {
		fmt.Println(subStrHash(s[i], power[i], modulo[i], k[i], hashValue[i]))
	}
}
