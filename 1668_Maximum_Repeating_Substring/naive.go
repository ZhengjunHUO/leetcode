package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/strings"
)

func maxRepeating(sequence string, word string) int {
	rslt := strings.NewPatternFinder(word).FindIn(sequence)
	if len(rslt) == 0 {
		return 0
	}

	max, curr := 1, 1

	for i := 1; i < len(rslt); i++ {
		if rslt[i] - rslt[i-1] == len(word) {
			curr++
			if curr > max {
				max = curr
			}
		}else{
			curr = 1
		}
	}

	return max
}

func main() {
	seqs := []string{"ababc", "ababc", "ababc", "dabckkabcabcabcdeabcabck"}
	words := []string{"ab", "ba", "ac", "abc"}

	for i := range seqs {
		fmt.Println(maxRepeating(seqs[i], words[i]))
	}
}
