package main

import (
	"fmt"
	"strconv"
)

// 思路同0190_Reverse_Bits，本题更加简单
func hammingWeight(num uint32) int {
	var cnt uint32

	for num > 0 {
		cnt += num & 1
		num = num >> 1
	} 
    
	return int(cnt)
}

func main() {
	n, _ := strconv.ParseUint("00000000000000000000000000001011", 2, 32)
	fmt.Println(hammingWeight(uint32(n)))

	n, _ = strconv.ParseUint("00000000000000000000000010000000", 2, 32)
	fmt.Println(hammingWeight(uint32(n)))

	n, _ = strconv.ParseUint("11111111111111111111111111111101", 2, 32)
	fmt.Println(hammingWeight(uint32(n)))

}
