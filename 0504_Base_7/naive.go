package main

import (
	"fmt"
)

func convertToBase7(num int) string {
	rslt := ""

	for num != 0 {
		rslt = fmt.Sprint(num%7) + rslt
		num /= 7
	}

	return rslt
}

func main() {
	nums := []int{100, -7}
	for i := range nums {
		fmt.Println(convertToBase7(nums[i]))
	}
}
