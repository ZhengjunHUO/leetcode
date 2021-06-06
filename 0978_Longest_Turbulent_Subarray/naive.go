package main

import (
	"fmt"
)

func maxTurbulenceSize(arr []int) int {
	asc, des, max, temp := 1, 1, 1, 0 
    
	for i:=1; i<len(arr); i++ {
		// 递增情况
		if arr[i] > arr[i-1] {
			// 核心所在，继承之前递减累计的长度（如果是连续递增则asc始终为2）
			asc = des + 1
			des = 1
		// 递减情况
		}else if arr[i] < arr[i-1]{
			// 同上
			des = asc + 1
			asc = 1
		}else{
			// 如果持平则重置计数
			asc, des = 1, 1
		}

		if asc > des {
			temp = asc
		}else{
			temp = des
		}
	
		if temp > max {
			max = temp
		}
	}

	return max
}

func main() {
	arr := []int{9,4,2,10,7,8,8,1,9}
	fmt.Println(maxTurbulenceSize(arr))
	
	arr = []int{4,8,12,16}
	fmt.Println(maxTurbulenceSize(arr))

	arr = []int{100}
	fmt.Println(maxTurbulenceSize(arr))
}
