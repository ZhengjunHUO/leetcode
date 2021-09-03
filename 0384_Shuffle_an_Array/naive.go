package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	origin		[]int
	shuffled	[]int
}

func Constructor(nums []int) Solution {
	origin, shuffled := make([]int, len(nums)), make([]int, len(nums))
	copy(origin, nums)
	copy(shuffled, nums)
	return Solution{origin, shuffled}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	copy(this.shuffled, this.origin)
	return this.shuffled
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.shuffled)
	// 共有n!中可能性
	for i:=0; i<n-1; i++ {
		r := this.random(i, n-1)
		this.shuffled[i], this.shuffled[r] = this.shuffled[r], this.shuffled[i]	
	}

	return this.shuffled
}

// return a random int in [a,b]
func (this *Solution) random(a, b int) int {
	return a + rand.Intn(b-a+1)
}

func main() {
	obj := Constructor([]int{1, 2, 3})
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Reset())

	dict := make(map[string]int)
	for i:=0; i<60000; i++ {
		dict[fmt.Sprintf("%v", obj.Shuffle())] += 1
		obj.Reset()
	}

	fmt.Println(dict)
}
