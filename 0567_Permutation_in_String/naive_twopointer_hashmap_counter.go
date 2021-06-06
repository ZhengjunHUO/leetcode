package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	m := make(map[int]int)
	for i := range s1 {
		m[int(s1[i])] += 1
	}    
	
	count := len(s1)

	lp, rp := 0, 0

	for rp < len(s2) {
		if m[int(s2[rp])] > 0 {
			count -= 1
		}
		m[int(s2[rp])] -= 1
		rp += 1

		for count==0 {
			if (rp-lp) == 2 {
				return true
			}

			m[int(s2[lp])] += 1
			if m[int(s2[lp])] > 0 {
				count += 1
			}
			lp += 1
		}
	}

	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1,s2))

	s2 = "eidboaoo"
	fmt.Println(checkInclusion(s1,s2))
}
