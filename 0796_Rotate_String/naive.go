package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func rotateString(s string, goal string) bool {
	pf := godtype.NewPatternFinder(goal)

	text := s+s
	
	return len(pf.FindIn(text)) != 0
}

func main() {
	ss := []string{"abcde", "abcde"}
	goals := []string{"cdeab" , "abced"}

	for i := range(ss) {
		fmt.Println(rotateString(ss[i], goals[i]))
	}
}
