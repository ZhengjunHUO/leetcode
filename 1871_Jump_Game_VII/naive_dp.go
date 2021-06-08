package main

import (
	"fmt"
)

// reach记录了所有元素是否可以到达，初始值0为不可到达，1为可以到达，最后返回列表最后一个元素是否大于0
// 当前元素是否可以到达只需要关心reach[i-maxJump, i-minJump]之间是否有1
func canReach(s string, minJump int, maxJump int) bool {
	// cnt 统计[i-maxJump, i-minJump]区间内有多少个1
	n, cnt := len(s), 0
	reach := make([]int, n)
	reach[0] = 1
	
	for i:=1; i<n; i++ {
		// 随着窗口的推进，计算新元素加入和旧元素退出对于cnt的影响
		if i>=minJump {
			cnt += reach[i-minJump]
		} 
		if i>maxJump{
			cnt -= reach[i-maxJump-1]
		}

		// 得出当前元素是否可以到达的结论
		if s[i] == '0' && cnt > 0 {
			reach[i] = 1
		}
	}	

	fmt.Println(reach)
	return reach[n-1] > 0
}

func main() {
	s := "011010"
	minJump, maxJump := 2, 3
	fmt.Println(canReach(s, minJump, maxJump))

	s = "01101110"
	minJump, maxJump = 2, 3
	fmt.Println(canReach(s, minJump, maxJump))
}
