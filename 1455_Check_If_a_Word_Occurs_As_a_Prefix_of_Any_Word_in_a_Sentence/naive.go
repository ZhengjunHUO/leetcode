package main

import (
	"fmt"
	"strings"
	"github.com/ZhengjunHUO/godtype"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	pf := godtype.NewPatternFinder(searchWord)

	str := strings.Split(sentence, " ")
	for i := range str {
		rslt := pf.FindIn(str[i])
		if len(rslt) > 0 && rslt[0] == 0 {
			return i+1
		}
	}

	return -1
}

func main() {
	sents := []string{"i love eating burger", "this problem is an easy problem", "i am tired", "i use triple pillow", "hello from the other side"}
	sws := []string{"burg", "pro", "you", "pill", "they"}

	for i := range sents {
		fmt.Println(isPrefixOfWord(sents[i], sws[i]))
	}
}
