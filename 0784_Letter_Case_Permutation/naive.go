package main

import (
	"fmt"
)

func backtrack(s string, rslt *[]string, curr string, currIdx int) {
	if len(curr) == len(s) {
		*rslt = append(*rslt, curr)
	}

	for i := currIdx; i < len(s); i++ {
		backtrack(s, rslt, curr + string(s[i]), i+1)

		switch {
		case s[i]>=65 && s[i]<=90:
			backtrack(s, rslt, curr + string(s[i]+32), i+1)
			break
		case s[i]>=97 && s[i]<=122:
			backtrack(s, rslt, curr + string(s[i]-32), i+1)
			break
		}
	} 

	
}

func letterCasePermutation(s string) []string {
	var rslt []string
	backtrack(s, &rslt, "", 0)
	return rslt 
}

func main() {
	str := []string{"a1b2", "3z4", "12345", "0"}	
	for _, v := range str {
		fmt.Println(letterCasePermutation(v))
	}
}
