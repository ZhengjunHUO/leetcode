package main

import (
	"fmt"
	"math"
	"strings"
)

/*
  字符串a*N，需要包含b，则其长度len(a*N)>=len(b)，考虑到如"abcd"和"cdabcdab"的情况
  需要再append一个额外的a才能完全包含
  strings.Index(s, subs)函数可以用0028_Implement_strStr中实现的算法替代
*/
func repeatedStringMatch(a string, b string) int {
	nb := int(math.Ceil(float64(len(b))/float64(len(a))))

	for i:=0; i<=1; i++ {
		if strings.Index(strings.Repeat(a, nb+i), b) != -1 {
			return nb+i
		}
	}

	return -1
}

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("a","aa"))
	fmt.Println(repeatedStringMatch("a","a"))
	fmt.Println(repeatedStringMatch("abc","wxyz"))
}
