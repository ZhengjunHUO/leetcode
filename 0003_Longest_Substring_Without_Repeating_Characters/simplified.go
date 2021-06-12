package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	dict := make(map[byte]int)	
	begin, maxVal := -1, 0
    
	for i := range s {
		if v, ok := dict[s[i]]; ok {
			begin = v
		}
		
		dict[s[i]] = i
	
		if temp := i - begin; temp > maxVal {
			maxVal = temp 
		}
	}	

	return maxVal
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}
