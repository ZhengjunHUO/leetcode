package main

import (
	"fmt"
)

type Deque struct {
	Elems []interface{}
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) pushFirst(elem interface{}) {
	temp := []interface{}{elem}
	d.Elems = append(temp, d.Elems...)
}

func (d *Deque) pushLast(elem interface{}) {
	d.Elems = append(d.Elems, elem)
}

func (d *Deque) popFirst() interface{} {
	if d.isEmpty() {
		return nil
	}

	first := d.Elems[0]
	d.Elems = d.Elems[1:]
	return first
}

func (d *Deque) popLast() interface{} {
	if d.isEmpty() {
		return nil
	}

	n := len(d.Elems)
	last := d.Elems[n-1]
	d.Elems = d.Elems[:n-1]
	return last
}

func (d *Deque) peekFirst() interface{} {
	if d.isEmpty() {
		return nil
	}

	return d.Elems[0]
}

func (d *Deque) peekLast() interface{} {
	if d.isEmpty() {
		return nil
	}

	return d.Elems[len(d.Elems)-1]
}

func (d *Deque) isEmpty() bool {
	if len(d.Elems) == 0 {
		return true
	}
	
	return false
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums) 

	// 新元素在尾部加入，头部为当前范围内的最大值
	// 保存的值为index
	dq := NewDeque()
	rslt := make([]int, n-k+1)
	rsltIdx := 0    

	for i := range nums{
		// 去除已经落在范围外的局部最大值（检查最大值的index）
		for (!dq.isEmpty()) && (dq.peekFirst().(int) < i - k + 1) {
			dq.popFirst()
		}	

		// 从尾向头去除小于当前nums[i]值的队列元素（检查index对应的值）
		for (!dq.isEmpty()) && (nums[dq.peekLast().(int)]) < nums[i] {
			dq.popLast()
		}

		// 将当前元素的index插入队尾
		dq.pushLast(i)
		if i >= k - 1 {
			// 头部为局部最大值的index，将对应数值保存至结果
			rslt[rsltIdx] = nums[dq.peekFirst().(int)]
			rsltIdx += 1
		}
	}

	return rslt	
}

func main() {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))

	nums = []int{1}
	k = 1
	fmt.Println(maxSlidingWindow(nums, k))

	nums = []int{1, -1}
	k = 1
	fmt.Println(maxSlidingWindow(nums, k))

	nums = []int{9, 11}
	k = 2
	fmt.Println(maxSlidingWindow(nums, k))

	nums = []int{4, -2}
	k = 2
	fmt.Println(maxSlidingWindow(nums, k))
}
