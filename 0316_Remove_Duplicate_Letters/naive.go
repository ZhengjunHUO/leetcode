package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func removeDuplicateLetters(s string) string {
	stack := godtype.NewStack()
	instack := make(map[byte]bool)
	counter := make(map[byte]int)
	for i := range s {
		counter[s[i]]++
	}

	for i := range s {
		if v, ok := instack[s[i]]; ok {
			if v == true {
				continue
			}
		}

		for !stack.IsEmpty() && s[i] < stack.Peek().(byte) {
			if counter[stack.Peek().(byte)] == 0 {
				break
			}

			instack[stack.Pop().(byte)] = false
		}

		stack.Push(s[i])
		counter[s[i]]--
		instack[s[i]] = true
	}	

	rslt := make([]byte, stack.Size())
	n := stack.Size() - 1
	for !stack.IsEmpty() {
		rslt[n] = stack.Pop().(byte)
		n--
	}

	return string(rslt)
}

func main() {
	s := []string{"bcabc", "cbacdcbc", "bcac"}
	for i := range s {
		fmt.Println(removeDuplicateLetters(s[i]))
	}
}
