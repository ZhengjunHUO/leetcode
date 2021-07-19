package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)

	defer func(){
		*h = (*h)[:n-1]	
	}()

	return (*h)[n-1]
}

type KthLargest struct {
	ih *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	temp := make(IntHeap, len(nums))
	for i, v := range nums {
		temp[i] = v
	}

	kl := KthLargest{
		ih: &temp,
	}

	heap.Init(kl.ih)

	for i:=0; i<len(nums)-3; i++ {
		heap.Pop(kl.ih)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.ih, val) 
	heap.Pop(this.ih)
	rslt := heap.Pop(this.ih)
	heap.Push(this.ih, rslt)
	return rslt.(int)
}

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2});
	val := []int{3, 5, 10, 9, 4}
	for _, v := range val {
		fmt.Println(obj.Add(v))
	}
}
