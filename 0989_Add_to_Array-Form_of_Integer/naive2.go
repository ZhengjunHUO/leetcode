package main

import (
	"fmt"
)

func addToArrayForm(num []int, k int) []int {
	for i:=len(num)-1; i>=0; i-- {
		temp := num[i] + k
		k, num[i] = temp/10, temp%10
	}

	for k > 0 {
		num = append([]int{k%10}, num...)
		k /= 10
	}

	return num
}

func main() {
	num := []int{1,2,0,0}
	k := 34
	fmt.Println(addToArrayForm(num, k))

	num = []int{2,7,4}
	k = 181
	fmt.Println(addToArrayForm(num, k))

	num = []int{2,1,5} 
	k = 806
	fmt.Println(addToArrayForm(num, k))

	num = []int{9,9,9,9,9,9,9,9,9,9}
	k = 1
	fmt.Println(addToArrayForm(num, k))
}
