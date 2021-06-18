package main

import (
	"fmt"
)

var start, max int

func searchBidirectional(s string, lp int, rp int) {
	for lp >= 0 && rp < len(s) && s[lp] == s[rp] {
		lp -= 1
		rp += 1
	}

	if tmp := rp - lp - 1; max < tmp {
		start = lp + 1
		max = tmp
	}
}

func longestPalindrome(s string) string {
	start, max = -1, 0

	// 每次以某个元素为中心，向两边确认是否对称
	for i := 0 ; i < len(s); i++ {
		// 假定中心为1个的情况
		searchBidirectional(s, i, i)
		// 假定中心为2个的情况
		searchBidirectional(s, i, i+1)
	}

	return s[start:start+max]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
	fmt.Println(longestPalindrome("fcabacdfgdcabaf"))
}
