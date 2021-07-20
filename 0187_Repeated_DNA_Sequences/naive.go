package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	dict := make(map[string]int)
	for i := 0; i + 9 < len(s); i++ {
		dict[s[i:i+10]] += 1
	}

	rslt := make([]string, 0, len(dict))
	for k, v := range dict {
		if v > 1 {
			rslt = append(rslt, k)
		}
	}

	return rslt
}

func main() {
	str := []string{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", "AAAAAAAAAAAAA"}
	for _,v := range str {
		fmt.Println(findRepeatedDnaSequences(v))
	}
}
