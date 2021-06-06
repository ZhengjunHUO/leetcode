package main

import (
	"fmt"
)

func jump(nums []int) int {
	// currJmpBorder当前这一跳可以覆盖的范围
	// currJmpReachMax该范围内的每个点出发可以到达的最远距离
	n, jmpCnt, currJmpBorder, currJmpReachMax := len(nums), 0, 0, 0 
	for i:=0; i<n-1; i++ {
		if temp := i + nums[i]; temp > currJmpReachMax {
			currJmpReachMax = temp
		}

		// 如果当前一跳范围内有结点可以直接跳到终点，则直接从该点起跳，算法结束
		if currJmpReachMax >= n-1 {
			jmpCnt+=1
			//fmt.Println("Already find the best solution at i =", i)
			break
		}

		// 下一个结点即将超出这一跳的覆盖范围，所以需要增加jmpCnt，边界更新为这一跳可到达的最远距离
		if i == currJmpBorder {
			jmpCnt+=1
			currJmpBorder = currJmpReachMax 
		}
	}

	return jmpCnt
}

func main() {
	nums := []int{2,3,1,1,4}
	fmt.Println(jump(nums))

	nums = []int{2,3,0,1,4}
	fmt.Println(jump(nums))
}
