package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

type MyStack struct {
	Queue	*godtype.Queue
	TopElem	int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{godtype.NewQueue(), 0}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Queue.Push(x)
	this.TopElem = x
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	size := this.Queue.Size()

	for size > 2 {
		this.Queue.Push(this.Queue.Pop().(int))
		size--
	}

	this.TopElem = this.Queue.Peek().(int)
	this.Queue.Push(this.Queue.Pop().(int))

	return this.Queue.Pop().(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.TopElem
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Queue.IsEmpty()
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
