package main

import (
	"fmt"
)

func largestOddNumber(num string) string {
	idx := -1
	for i:=len(num)-1; i>=0; i-- {
		if (num[i] - '0') % 2 == 1 {
			idx = i
			break
		}
	}
	
	if idx == -1 {
		return ""
	}

	return num[:idx+1]
}

func main() {
	nums := []string{"52", "4206", "35427"}
	for i := range nums {
		fmt.Println(largestOddNumber(nums[i]))
	}
}
