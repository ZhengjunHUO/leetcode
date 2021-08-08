package main

import (
	"fmt"
)

func reverseString(s []byte) {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	ss := [][]byte{[]byte{'h','e','l','l','o'}, []byte{'H','a','n','n','a','h'}}
	for i := range ss {
		reverseString(ss[i])
		fmt.Printf("%s\n", string(ss[i]))
	}
}
