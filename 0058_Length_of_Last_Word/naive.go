package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	i, rslt := len(s) - 1, 0

	// 移除末尾可能存在的空格
	for ; i >= 0; i-- {
		if s[i] != 32 {
			break
		}
	}

	// 统计最后一个词的长度
	for ; i >= 0; i-- {
		if s[i] == 32 {
			break
		}
		rslt++ 
	}

	return rslt
}

func main() {
	str := []string{"Hello World ", " "}
	for i := range str {
		fmt.Println(lengthOfLastWord(str[i]))
	}
}
