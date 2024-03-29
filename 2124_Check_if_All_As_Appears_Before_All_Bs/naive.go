package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/strings"
)

// 只需要判断是否含有"ba"这个pattern
func checkString(s string) bool {
	pf := strings.NewPatternFinder("ba")
	rslt := pf.FindIn(s)

	return len(rslt) == 0
}

func main() {
	s := []string{"aaabbb", "abab", "bbb"}
	for i := range s {
		fmt.Println(checkString(s[i]))
	}
}
