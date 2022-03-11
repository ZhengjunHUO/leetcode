package main

import (
	"fmt"
)

func reversePrefix(word string, ch byte) string {
	l, r := 0, 0
	ret := []byte(word)
	for i := range ret {
		if ret[i] == ch {
			r = i
			break
		}
	}

	for l < r {
		ret[l], ret[r] = ret[r], ret[l]
		l++
		r--
	}

	return string(ret)
}

func main() {
	word := []string{"abcdefd", "xyxzxe", "abcd"}
	ch := []byte{'d', 'z', 'z'}

	for i := range word {
		fmt.Println(reversePrefix(word[i], ch[i]))
	}
}
