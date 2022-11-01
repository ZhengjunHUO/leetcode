package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func smallestSubsequence(s string) string {
	stack := datastruct.NewStack([]byte{})
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

		for !stack.IsEmpty() && s[i] < stack.Peek() {
			if counter[stack.Peek()] == 0 {
				break
			}

			instack[stack.Pop()] = false
		}

		stack.Push(s[i])
		counter[s[i]]--
		instack[s[i]] = true
	}

	rslt := make([]byte, stack.Size())
	n := stack.Size() - 1
	for !stack.IsEmpty() {
		rslt[n] = stack.Pop()
		n--
	}

	return string(rslt)
}

func main() {
	s := []string{"bcabc", "cbacdcbc", "bcac"}
	for i := range s {
		fmt.Println(smallestSubsequence(s[i]))
	}
}
