package main

import (
	"fmt"
)

// @lee215的递归更简洁，采用他的方案替代自己的 
func canReach(arr []int, start int) bool {
	if start >=0 && start <len(arr)-1 && arr[start]>=0 {
		arr[start] = -arr[start]
		return arr[start] == 0 || canReach(arr, start+arr[start]) || canReach(arr, start-arr[start]) 
	}  
	return false
}

func main() {
	arr := []int{4,2,3,0,3,1,2}
	start := 5
	fmt.Println(canReach(arr, start))

	arr = []int{4,2,3,0,3,1,2}
	start = 0
	fmt.Println(canReach(arr, start))

	arr = []int{3,0,2,1,2}
	start = 2
	fmt.Println(canReach(arr, start))
}
