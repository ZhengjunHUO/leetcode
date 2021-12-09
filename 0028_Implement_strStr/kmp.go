package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return 0
	}

	pf := godtype.NewPatternFinder(needle)
	rslt := pf.FindIn(haystack)
	if len(rslt) == 0 {
		return -1
	} 

	return rslt[0]
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
}
