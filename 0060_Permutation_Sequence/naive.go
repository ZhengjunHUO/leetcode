package main

import (
	"fmt"
	"strings"
)

// 计算某数的阶乘
func factorial(n int) int {
	rslt := 1
	for n > 0 {
		rslt *= n
		n -= 1
	}

	return rslt
}

func getPerm(n int, k int, curr []int, rslt *string) {
	// 找第一个perm即最小的组合，只需要将剩余的数列curr(递增)append到结果并返回
	if k <= 1 {
		*rslt += strings.Trim(strings.Join(strings.Split(fmt.Sprint(curr), " "), ""), "[]")		
		return
	}

	// 当前n的所有排列组合数，除以n，可以得到以n中每个数开头的组合数
	nbBeginWithN := factorial(n)/n
	// 再结合k来确定其落在以哪个数开头的perm集合中，将该数append到结果中
	currIdx := k/nbBeginWithN 
	*rslt += fmt.Sprint(curr[currIdx])

	// 在原数列中剔除该数但保留其递增属性
	getPerm(n-1, k%nbBeginWithN, append(curr[0:currIdx], curr[currIdx+1:len(curr)]...), rslt)
}

func getPermutation(n int, k int) string {
	curr := make([]int, n)	
	for i:=0; i<n; i++ {
		curr[i] = i + 1
	}

	rslt := ""
	getPerm(n, k, curr, &rslt)
	return rslt
}

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
}
