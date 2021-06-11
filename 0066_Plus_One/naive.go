package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := 1

	for i:=len(digits)-1; i>=0; i-- {
		digits[i] += carry
		carry = digits[i]/10
		if carry == 0 {
			return digits
		}
		digits[i] %= 10
	} 

	return append([]int{carry}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{1,2,3}))
	fmt.Println(plusOne([]int{4,3,2,1}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{9,9,9}))
}
