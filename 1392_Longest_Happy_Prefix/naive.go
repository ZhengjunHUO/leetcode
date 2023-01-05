package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/strings"
)

func longestPrefix(s string) string {
	lps := strings.NewPatternFinder(s).Lps

	return s[:lps[len(lps)-1]]
}

func main() {
	ss := []string{"level", "ababab", "leetcodeleet", "a"}
	for i := range ss {
		fmt.Println(longestPrefix(ss[i]))
	}
}
