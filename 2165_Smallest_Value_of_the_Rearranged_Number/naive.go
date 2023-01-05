package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func smallestNumber(num int64) int64 {
	if num == 0 {
		return 0
	}

	var ret int64
	var zeroCount int
	var pq *datastruct.PriorityQueue[int, int]
	var isPositive bool

	if num < 0 {
		// for negative number we use a max priority queue
		pq = datastruct.NewPQ([]int{}, []int{}, false)
		isPositive = false
		num *= -1
	}else{
		// for positive number we use a min priority queue
		pq = datastruct.NewPQ([]int{}, []int{}, true)
		isPositive = true
	}

	for num > 0 {
		if temp := int(num%10) ; temp == 0 {
			zeroCount++
		}else{
			// push digit to priority queue
			pq.Push(temp, temp)
		}

		num /= 10
	}

	// decide the leading digit
	ret = int64(pq.Pop())
	if isPositive {
		// rebuild number by assembling digits in ascending order, begin with zero(s)
		for zeroCount > 0 {
			ret *= 10
			zeroCount--
		}

		for pq.Size() > 0 {
			ret = ret*10 + int64(pq.Pop())
		}
	}else{
		// rebuild number by assembling digits in descending order, zero(s) come at last
		for pq.Size() > 0 {
			ret = ret*10 + int64(pq.Pop())
		}

		for zeroCount > 0 {
			ret *= 10
			zeroCount--
		}

		ret *= -1
	}

	return ret
}

func main() {
	nums := []int64{310,-7605}
	for i := range nums {
		fmt.Println(smallestNumber(nums[i]))
	}
}
