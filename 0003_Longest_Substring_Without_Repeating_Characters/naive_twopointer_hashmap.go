package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dict := make(map[byte]int)	
	lp, rp, currVal, maxVal := 0, 0, 0 ,0
    
	for rp < n {
		curr := s[rp]
		if _, ok := dict[curr]; ok {
			for lp < rp {
				delete(dict, s[lp])
				currVal -= 1
				lp += 1 

				if s[lp-1] == curr {
					break
				}
			}	
		}

		dict[curr] = rp
		currVal += 1
	
		if currVal > maxVal {
			maxVal = currVal
		}
		rp += 1
	}	

	return maxVal
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}
