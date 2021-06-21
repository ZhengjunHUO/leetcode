package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// 带负号的和末位为0的数（0本身除外）都不是回文
	if x < 0 || ( x !=0 && x%10 == 0) {
		return false
	}

	val := 0

	/* 
	  剥离末位的数倒序构建另一个整型，直到该数等于或大于自己
	  (1) 等于的情况说明原数的位数是偶数且是回文，返回true
          (2) 大于的话，如果原数的位数是偶数则不是回文(eg. 1231 => 13 > 12)
	      如果原数的位数是奇数，则生成的数的位数多一位，需要移除后再比较(eg. 12321 => 123 > 12 )
	*/
	for x > val {
		val = val * 10 + x % 10
		x /= 10
	}
	
	return x == val || x == val/10
}

func main() {
	test := []int{12321, 6776, 7776, -101, 100, 55, 76, 6}
	for _,v := range test {
		fmt.Println(isPalindrome(v))
	} 
}
