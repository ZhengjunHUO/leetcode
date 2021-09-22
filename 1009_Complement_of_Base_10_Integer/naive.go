package main

import (
	"fmt"
)

// 重点是需要找出n开头有几个零，然后用n二进制长度的全1去XOR即可
func bitwiseComplement(n uint32) uint32 {
	// 处理0的情况
	if n == 0 {
		return 1
	}

	var mask uint32 = 0xFFFFFFFF
	// mask中尾部的0的数量即是n二进制表示的长度
	for (n & mask) > 0 {
		mask = mask << 1
	}

	// 反转mask得到适应n长度的全1
	return ^mask ^ n
}

func main() {
	nums := []uint32{5,7,10,0}
	for i := range nums {
		fmt.Println(bitwiseComplement(nums[i]))
	}
}
