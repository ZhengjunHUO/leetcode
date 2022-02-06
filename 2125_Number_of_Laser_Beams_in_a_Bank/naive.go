package main

import (
	"fmt"
)

func numberOfBeams(bank []string) int {
	var c, former, ret int

	for i := range bank {
		// 统计每一行上有几个设备
		c = 0
		for j := range bank[i] {
			if bank[i][j] == '1' {
				c++
			}
		}

		if c != 0 {
			// 如果前排有设备则计算并增加镭射数量
			if former != 0 {
				ret += former * c
			}
			// 更新前排设备数
			former = c
		}
		// 如果没有设备则直接跳过
	}

	return ret
}

func main() {
	banks := [][]string{[]string{"011001","000000","010100","001000"}, []string{"000","111","000"}}
	for i := range banks {
		fmt.Println(numberOfBeams(banks[i]))
	}
}
