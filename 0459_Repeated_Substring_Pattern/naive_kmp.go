package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	tab := gtoolkit.GetKmpTable(s)    

	/* 
	  符合要求的s的kmp表的最后一格的内容为len(s)-len(pattern)，如
	    a b a a b a a b a
	    0 0 1 1 2 3 4 5 6
	*/
	return tab[n-1] > 0 && n%(n-tab[n-1]) == 0
}

func main() {
	fmt.Println("abab: ", repeatedSubstringPattern("abab"))
	fmt.Println("aba: ", repeatedSubstringPattern("aba"))
	fmt.Println("abcabcabcabc: ", repeatedSubstringPattern("abcabcabcabc"))
}
