package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	strMap := make(map[string]int) 
	for _,v := range words {
		strMap[v] += 1
	}
	count := len(words)

	lp, rp, step := 0, 0, len(words[0])
	rslt := []int{}

	for rp < len(s) {
		if strMap[s[rp:rp+step]] > 0 {
			count -= 1
		}
		strMap[s[rp:rp+step]] -= 1
		rp += step
		
		for count==0 {
			if (rp - lp) == len(words) * step {
				rslt = append(rslt, lp)
			}

			strMap[s[lp:lp+step]] += 1
			if strMap[s[lp:lp+step]] > 0 {
				count += 1
			}

			lp = lp + step
		}
	}
	return rslt
}



func main() {
	s := "barfoothefoobarman"
	words := []string{"foo","bar"}
	fmt.Println(findSubstring(s, words))

	s = "wordgoodgoodgoodbestword"
	words = []string{"word","good","best","word"}
	fmt.Println(findSubstring(s, words))

	s = "barfoofoobarthefoobarman"
	words = []string{"bar","foo","the"}
	fmt.Println(findSubstring(s, words))
}
