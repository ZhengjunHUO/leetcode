package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	// 左括号累计出现次数，右括号累计出现次数，历史最大合法括号对长度
	lcum, rcum, max := 0, 0, 0
	    
	// 从左到右遍历字符串
	for i := 0 ; i < len(s) ; i++ {
		switch s[i] {
		case '(':
			lcum += 1
		case ')':
			rcum += 1
		}

		// 左括号和右括号成对，记录截止目前得到的最大长度
		if lcum == rcum {
			if temp := 2 * rcum; temp > max {
				max = temp
			}
		}

		// 不合法，重置计数器
		if rcum > lcum {
			lcum, rcum = 0, 0
		}

		// 无法判断lcum > rcum下的情况，需要反方向再遍历一遍
	}

	lcum, rcum = 0, 0 

	// 再从右到左遍历字符串，因为上一轮遍历没能在max中记录在lcum>rcum时的合法长度
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			lcum += 1
		case ')':
			rcum += 1
		}

		if lcum == rcum {
			if temp := 2 * lcum; temp > max {
				max = temp
			}
		}
		
		if lcum > rcum {
			lcum, rcum = 0, 0 
		}
	}

	return max
}

func main() {
	test := []string{"(()", ")()())", "", "()((()))((())", "((()())"}
	for _,v := range test {
		fmt.Printf("%s: %v\n", v, longestValidParentheses(v))
	}
}
