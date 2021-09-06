package main

import (
	"fmt"
)

func minAddToMakeValid(s string) int {
	leftCnt, rslt := 0, 0

	for i := range s {
		if s[i] == byte('(') {
			leftCnt++
		}else{
			if leftCnt == 0 {
				rslt++
			}else{
				leftCnt--
			}
		}
	}

	return rslt + leftCnt	
}

func main() {
	ss := []string{"())", "(((", "()", "()))(("}
	for i := range ss {
		fmt.Println(minAddToMakeValid(ss[i]))
	}
}
