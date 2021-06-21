package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	// 如果s是循环的字符串，那么s+s组合的字符串去掉首尾各一个字符后还是可以找到1个s
	// strings.Index(s, subs)函数可以用0028_Implement_strStr中实现的算法替代
	return strings.Index((s+s)[1:2*len(s)-1], s) != -1
}

func main() {
	fmt.Println("abab: ", repeatedSubstringPattern("abab"))
	fmt.Println("aba: ", repeatedSubstringPattern("aba"))
	fmt.Println("abcabcabcabc: ", repeatedSubstringPattern("abcabcabcabc"))
}
