package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

type MyQueue struct {
	sOrigin		*godtype.Stack
	sInverse	*godtype.Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{godtype.NewStack(), godtype.NewStack()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.sOrigin.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Peek() == -1 {
		return -1
	}

	return this.sInverse.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.sInverse.IsEmpty() {
		for ! this.sOrigin.IsEmpty() {
			this.sInverse.Push(this.sOrigin.Pop())
		}
	}

	if this.Empty() {
		return -1
	}

	return this.sInverse.Peek().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.sOrigin.IsEmpty() && this.sInverse.IsEmpty()
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
