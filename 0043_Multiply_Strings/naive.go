package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	l1, l2, v1, sum := len(num1), len(num2), 0, 0
	// 两个位数为n, m的数的乘积位数不会超过n+m
	prod := make([]int, l1+l2)    
	rslt := ""
	
	// 从低位开始按位相乘，每次相乘所得的积在prod的中的位置为i+j（高位）, i+j+1（低位）
	for i:=l1-1; i>=0; i-- {
		v1 = int(num1[i] - '0')	
		for j:=l2-1; j>=0; j-- {
			// 重要：需要将对应的低位置的值加入计算进位，而不是直接+=上去
			sum = v1 * int(num2[j] - '0') + prod[i+j+1]
			prod[i+j] += sum/10
			prod[i+j+1] = sum%10 
		}
	}

	// 跳过prod中开始可能存在的0
	begin := 0
	for begin < len(prod)-1 {
		if prod[begin] != 0 {
			break
		}
		begin += 1	
	}

	for ;begin < len(prod); begin++ {
		rslt = fmt.Sprintf("%s%v", rslt, prod[begin])
	}
	
	return rslt
}

func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("99", "99"))
}
