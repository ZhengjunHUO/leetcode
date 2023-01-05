package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/strings"
)

func stringMatching(words []string) []string {
	s := ""
	rslt := []string{}

	for i := range words {
		s = s + words[i] + " "
	}

	for i := range words {
		if len(strings.NewPatternFinder(words[i]).FindIn(s)) >= 2 {
			rslt = append(rslt, words[i])
		}
	}

	return rslt
}

func main() {
	words := [][]string{[]string{"mass","as","hero","superhero"}, []string{"leetcode","et","code"}, []string{"blue","green","bu"}}
	for i := range words {
		fmt.Println(stringMatching(words[i]))
	}
}
