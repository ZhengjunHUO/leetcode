package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func addToArrayForm(num []int, k int) []int {
	n := len(num)
	i, v1, v2, carry := n-1, 0, 0, 0

	q := datastruct.NewQueue([]int{})
	for k >= 10 {
		q.Push(k%10)
		k = k/10
	}
	q.Push(k)

	rslt := []int{}

	for i>=0 || !q.IsEmpty() || carry > 0 {
		if i>=0 {
			v1 = num[i]
			i-=1
		}else{
			v1 = 0
		}

		if !q.IsEmpty() {
			v2 = q.Pop()
		}else{
			v2 = 0
		}

		sum := v1 + v2 + carry
		carry = sum/10

		rslt = append([]int{sum%10}, rslt...)
	}

	return rslt
}

func main() {
	num := []int{1,2,0,0}
	k := 34
	fmt.Println(addToArrayForm(num, k))

	num = []int{2,7,4}
	k = 181
	fmt.Println(addToArrayForm(num, k))

	num = []int{2,1,5}
	k = 806
	fmt.Println(addToArrayForm(num, k))

	num = []int{9,9,9,9,9,9,9,9,9,9}
	k = 1
	fmt.Println(addToArrayForm(num, k))
}
