package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	m := make(map[int]int)
	for i := range p {
		m[int(p[i])] += 1
	}
	count := len(p)

	lp, rp := 0, 0
	rslt := []int{}

	for rp < len(s) {
		if m[int(s[rp])] > 0 {
			count -= 1
		}
		m[int(s[rp])] -= 1
		rp += 1

		for count == 0 {
			if (rp-lp) == len(p) {
				rslt = append(rslt, lp)
			}

			m[int(s[lp])] += 1
			if m[int(s[lp])] > 0 {
				count += 1
			}
			lp += 1
		}
	}

	return rslt
}

func main() {
	s, p := "cbaebabacd", "abc"
	fmt.Println(findAnagrams(s, p))

	s, p = "abab", "ab"
	fmt.Println(findAnagrams(s, p))
}
