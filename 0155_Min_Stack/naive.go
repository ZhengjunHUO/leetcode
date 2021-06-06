package main

import (
	"fmt"
)

// Min存储着：不考虑当前元素，本stack的最小值
// 用于被pop后更新整个stack的最小值
type StackElem struct {
	Val	int
	Min	int
}

// Min: 当前整个stack的最小值
type MinStack struct {
	Min	int
	Elems	[]StackElem
}

func Constructor() *MinStack {
	return &MinStack{
		Min: 2147483647,
		Elems: []StackElem{},
	}
}


func (this *MinStack) Push(val int)  {
	elem := StackElem{val, this.Min }

	if val < this.Min {
		this.Min = val		
	} 

	this.Elems = append(this.Elems, elem)
}


func (this *MinStack) Pop()  {
	this.Min = this.Elems[len(this.Elems)-1].Min
	this.Elems = this.Elems[:len(this.Elems)-1]
}


func (this *MinStack) Top() int {
	return this.Elems[len(this.Elems)-1].Val 
}


func (this *MinStack) GetMin() int {
	return this.Min 
}



func main() {
	s := Constructor();
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
}
