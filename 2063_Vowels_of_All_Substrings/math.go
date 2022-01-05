package main

import (
	"fmt"
)

/*
  从每一个字母(index i)的角度出发，它可能出现在word[l:r]中当
  0<=l<=i且i<=r<n，符合条件的l有i-0+1种，r有n-1-i+1种
*/
func countVowels(word string) int64 {
	dict := map[byte]bool{
		'a' : true,
		'e' : true,
		'i' : true,
		'o' : true,
		'u' : true,
	}

	n := len(word)
	var rslt int64
	for i:=0; i<n; i++ {
		if _, ok := dict[word[i]]; ok {
			rslt += int64((i+1) * (n-i))
		}
	}

	return rslt
}

func main() {
	word := []string{"aba", "abc", "ltcd"}
	for i := range word {
		fmt.Println(countVowels(word[i]))
	}
}
