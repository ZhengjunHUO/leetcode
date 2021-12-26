package main

import (
	"fmt"
)

var pattern = []int{0,1,1,2}

// 返回数列每个元素加上一个int的新数列
func ArrayPlusInt(arr []int, x int) []int {
	rslt := make([]int, len(arr))	
	for i := range arr {
		rslt[i] = arr[i] + x
	}

	return rslt
}

/* 
  每个迭代以4^n为分界线，内部可以分成4组，其中上个迭代4^(n-1)的结果为本次迭代的第一组 
  剩下的4^n - 4^(n-1)可以划分为3组，其值分别为第一组的+1, +1, +2
  如4^1为[0,1,1,2], 4^2在其基础上append [...]+1, [...]+1, [...]+2即可
*/
func countBits(n int) []int {
	rslt := []int{0,1,1,2}
	prod := 4
	
	for n >= prod {
		toAppend := [][]int{ArrayPlusInt(rslt, 1)}
		if n >= prod+len(rslt) {
			toAppend = append(toAppend, ArrayPlusInt(rslt, 1))
			if n >= prod+len(rslt)*2 {
				toAppend = append(toAppend, ArrayPlusInt(rslt, 2))
			}
		}

		for i := range toAppend {
			rslt = append(rslt, toAppend[i]...)
		}
		
		/*
		secondpart := ArrayPlusInt(rslt, 1)	
		thirdpart := ArrayPlusInt(rslt, 1)	
		fourthpart := ArrayPlusInt(rslt, 2)	
		rslt = append(rslt, secondpart...)
		rslt = append(rslt, thirdpart...)
		rslt = append(rslt, fourthpart...)
		*/

		prod *= 4
	}

	//fmt.Println(rslt)
	return rslt[:n+1]
}

func main() {
	n := []int{2,5,46}
	for i := range n {
		fmt.Println(countBits(n[i]))
	}
}
