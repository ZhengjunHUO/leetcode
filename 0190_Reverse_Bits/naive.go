package main

import (
	"fmt"
)

// 思路和0007_Reverse_Integer类似，时间复杂度是O(1)因为长度固定为32
func reverseBits(num uint32) uint32 {
	var rslt uint32
	move := 31

	for num > 0 {
		rslt += (num & 1) << move
		num = num >> 1
		move -= 1
	}

	return rslt
}

func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(4294967293))
}
