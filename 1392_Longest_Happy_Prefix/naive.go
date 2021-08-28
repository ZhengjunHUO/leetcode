package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func longestPrefix(s string) string {
	lps := godtype.NewPatternFinder(s).GetLPS()

	return s[:lps[len(lps)-1]]
}

func main() {
	ss := []string{"level", "ababab", "leetcodeleet", "a"}
	for i := range ss {
		fmt.Println(longestPrefix(ss[i]))
	}
}
