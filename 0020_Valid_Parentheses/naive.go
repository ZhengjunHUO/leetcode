package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func isValid(s string) bool {
	stack := datastruct.NewStack([]byte{})
	dict := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0 ; i < len(s) ; i++ {
		switch s[i] {
		case '(', '[', '{':
			stack.Push(s[i])
		case ')', ']', '}':
			if stack.IsEmpty() || stack.Pop() != dict[s[i]] {
				return false
			}
		}
	}

	if stack.Size() != 0 {
		return false
	}

	return true
}

func main() {
	test := []string{"()", "()[]{}", "(]", "([)]", "{[]}", "((())"}
	for _,v := range test {
		fmt.Printf("%s: %v\n", v, isValid(v))
	}
}
