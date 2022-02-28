package main

import (
	"fmt"
)

func countOperations(num1 int, num2 int) int {
	var ret int
	// until one goes to zero
	for num1 != 0 && num2 != 0 {
		// make num1 always bigger than num2 
		if (num1 < num2) {
			num1, num2 = num2, num1
		}
		// simulate times of substraction from the num1
		ret += num1/num2
		num1 = num1%num2
	}

	return ret
}

func main() {
	num1 := []int{2, 10, 3}
	num2 := []int{3, 10, 13}
	for i := range num1 {
		fmt.Println(countOperations(num1[i], num2[i]))
	}
}
