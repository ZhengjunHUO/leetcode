package main

import (
	"fmt"
)

func minInsertions(s string) int {
	waitR, rslt := 0, 0

	for i := range s {
		if s[i] == byte('(') {
			waitR += 2
			// 前一个左括号还需要补上一个右括号
			if waitR%2 != 0 {
				rslt++
				waitR--
			}
		}else{
			waitR--
			// 多余了一个右括号，需要补上一个左括号，同时需要后面再来一个右括号
			if waitR == -1 {
				rslt++
				waitR = 1
			}
			
		}
	}

	return rslt + waitR
}

func main() {
	ss := []string{"(()))", "())", "))())(", "((((((", ")))))))", ")("}
	for i := range ss {
		fmt.Println(minInsertions(ss[i]))
	}
}
