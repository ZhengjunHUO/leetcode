package main

import (
	"fmt"
	"strconv"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func kthLargestNumber(nums []string, k int) string {
	pq := datastruct.NewPQ([]int{}, []int{}, true)

	for i := range nums {
		x, _ := strconv.Atoi(nums[i])
		pq.Push(x, x)

		if pq.Size() > k {
			pq.Pop()
		}
	}

	return strconv.Itoa(pq.Peek())
}

func main() {
	nums := [][]string{[]string{"3","6","7","10"}, []string{"2","21","12","1"}, []string{"0","0"}}
	k := []int{4, 3, 2}
	for i := range nums {
		fmt.Println(kthLargestNumber(nums[i], k[i]))
	}
}
