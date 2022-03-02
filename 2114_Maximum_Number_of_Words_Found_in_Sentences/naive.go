package main

import (
	"fmt"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func mostWordsFound(sentences []string) int {
	var ret int
	for i := range sentences {
		ret = max(ret, len(strings.Split(sentences[i], " ")))
	}

	return ret
}

func main() {
	sentences := [][]string{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, []string{"please wait", "continue to fight", "continue to win"}}
	for i := range sentences {
		fmt.Println(mostWordsFound(sentences[i]))
	}
}
