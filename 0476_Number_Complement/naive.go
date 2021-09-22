package main

import (
	"fmt"
)

// 重点是需要找出num开头有几个零，然后用num二进制长度的全1去XOR即可
func findComplement(num uint32) uint32 {
	var mask uint32 = 0xFFFFFFFF
	// mask中尾部的0的数量即是num二进制表示的长度
	for (num & mask) > 0 {
		mask = mask << 1
	}

	// 反转mask得到适应num长度的全1
	return ^mask ^ num
}

func main() {
	nums := []uint32{5,1,13}
	for i := range nums {
		fmt.Println(findComplement(nums[i]))
	}
}
