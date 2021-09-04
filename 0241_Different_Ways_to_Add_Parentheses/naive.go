package main

import (
	"fmt"
	"strconv"
)

func calculate(op byte, a, b []int) []int {
	dict := make(map[int]bool)

	for i := range a {
		for j := range b {
			if op == byte('+') {
				dict[a[i]+b[j]]	= true
			}
			if op == byte('-') {
				dict[a[i]-b[j]]	= true
			}
			if op == byte('*') {
				dict[a[i]*b[j]]	= true
			}
		}
	}

	rslt := []int{}
	for k := range dict {
		rslt = append(rslt, k)
	}

	return rslt
}

// 复杂度为Catalan数列
func diffWaysToCompute(expression string) []int {
	rslt := []int{}

	for i:=0; i<len(expression); i++ {
		if expression[i] == byte('+') || expression[i] == byte('-') || expression[i] == byte('*') {
			rslt = append(rslt, calculate(expression[i], diffWaysToCompute(expression[:i]), diffWaysToCompute(expression[i+1:]))...)
		}
	}

	if len(rslt) == 0 {
		if temp, err := strconv.Atoi(expression); err == nil {
			return []int{temp}
		}
	}

	return rslt
}

func main() {
	exps := []string{"2-1-1", "2*3-4*5", "20-1-1"}
	for i := range exps {
		fmt.Println(diffWaysToCompute(exps[i]))
	}
}
