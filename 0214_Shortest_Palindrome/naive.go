package main

import (
	"fmt"
)

var start, max int

// 其实要找的是从字符串头开始最长的回文组，对0005的算法稍加改动即可
// 但不是最优解，做了很多无用功
func searchBidirectional(s string, lp int, rp int) {
	for lp >= 0 && rp < len(s) && s[lp] == s[rp] {
		lp -= 1
		rp += 1
	}

	if tmp := rp - lp - 1; lp == -1 && max < tmp {
		start = lp + 1
		max = tmp
	}
}

func shortestPalindrome(s string) string {
	start, max = -1, 0

	for i := 0 ; i < len(s); i++ {
		searchBidirectional(s, i, i)
		searchBidirectional(s, i, i+1)
	}

	rslt := s[start:]
	rp := start + max
	for rp < len(s) {
		rslt = s[rp:rp+1] + rslt	
		rp += 1
	}

	return rslt
  
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
	fmt.Println(shortestPalindrome("jkbacabdf"))
}
