package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
	"github.com/ZhengjunHUO/goutil/strings"
)

func canChoose(groups [][]int, nums []int) bool {
	// 将array转为字符串
	target := gtoolkit.IntSliceToString(nums)
	// 存储前一个subarray的在array中的index和其长度
	var prevLocation []int
	prevLen := 0

	for i := range groups {
		// 将subarray转为字符串
		pattern := gtoolkit.IntSliceToString(groups[i])
		psize := len(pattern)

		// 通过kmp算法找到其index
		pf := strings.NewPatternFinder(pattern)
		location := pf.FindIn(target)

		// 判断当前subarray是否在前一个subarray之后
		if len(prevLocation) > 0 && location[0] < (prevLocation[0] + prevLen) {
			return false
		}

		prevLocation, prevLen = location, psize
	}

	return true
}

func main() {
	groups := [][][]int{[][]int{[]int{1,-1,-1},[]int{3,-2,0}}, [][]int{[]int{10,-2},[]int{1,2,3,4}}, [][]int{[]int{1,2,3},[]int{3,4}}}
	nums := [][]int{[]int{1,-1,0,1,-1,-1,3,-2,0}, []int{1,2,3,4,10,-2}, []int{7,7,1,2,3,4,7,7}}

	for i := range nums {
		fmt.Println(canChoose(groups[i], nums[i]))
	}
}
