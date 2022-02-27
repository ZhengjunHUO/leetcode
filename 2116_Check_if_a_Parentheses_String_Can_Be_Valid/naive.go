package main

import (
	"fmt"
)

func canBeValid(s string, locked string) bool {
	// 不成对返回false
	if len(s)%2 != 0 {
		return false
	}

	var unlocked, left, right int
	// 从左到右判断每个不能修改的')'的都至少有一个固定'('
	// 或一个可修改的括号对应
	for i := range s {
		if locked[i] == '0' {
			unlocked++
		}else if s[i] == '(' {
			left++
		}else{
			right++
		}

		if unlocked + left < right {
			return false
		}
	}

	unlocked, left, right = 0, 0, 0
	// 同理从右到左观察固定'('的对应情况
	for i:=len(s)-1; i>=0; i-- {
		if locked[i] == '0' {
			unlocked++
		}else if s[i] == '(' {
			left++
		}else{
			right++
		}

		if unlocked + right < left {
			return false
		}
	}

	return true
}

func main() {
	s := []string{"))()))", "()()", ")"}
	locked := []string{"010100", "0000", "0"}
	for i := range s {
		fmt.Println(canBeValid(s[i], locked[i]))
	}
}
