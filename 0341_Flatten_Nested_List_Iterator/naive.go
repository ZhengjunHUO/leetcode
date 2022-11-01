package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
	"reflect"
)

type NestedInteger struct {
	val		int
	list		[]*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return len(this.list) == 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return this.val
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	this.list = append(this.list, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return this.list
}

func NewNestedInteger(list []interface{}) []*NestedInteger {
	rslt := []*NestedInteger{}

	for i:=0; i<len(list); i++ {
		if reflect.ValueOf(list[i]).Kind() == reflect.Int {
			rslt = append(rslt, &NestedInteger{list[i].(int), []*NestedInteger{}})
		}else{
			rslt = append(rslt, &NestedInteger{0, NewNestedInteger(list[i].([]interface{}))})
		}
	}

	return rslt
}


type NestedIterator struct {
	stack	*datastruct.Stack[*NestedInteger]
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	ni := &NestedIterator{datastruct.NewStack([]*NestedInteger{})}
	for i:=len(nestedList)-1; i>=0; i-- {
		ni.stack.Push(nestedList[i])
	}

	return ni
}

func (this *NestedIterator) Next() int {
	return this.stack.Pop().GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for ! this.stack.IsEmpty() {
		curr := this.stack.Peek()
		if curr.IsInteger() {
			return true
		}

		this.stack.Pop()
		list := curr.GetList()
		for i:=len(list)-1; i>=0; i-- {
			this.stack.Push(list[i])
		}
	}

	return false
}

func main() {
	//origList := []interface{}{[]interface{}{1,1},2,[]interface{}{1,1}}
	origList := []interface{}{1,[]interface{}{4,[]interface{}{6}}}
	ni := Constructor(NewNestedInteger(origList))
	rslt := []int{}

	for ni.HasNext() {
		rslt = append(rslt, ni.Next())
	}

	fmt.Println(rslt)
}
