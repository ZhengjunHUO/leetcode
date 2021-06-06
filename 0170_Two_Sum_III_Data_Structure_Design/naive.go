package main

import (
	"fmt"
)

type iTwoSum interface {
	AddInt(i int)
	Find(v int) bool
}

type TwoSum struct {
	list []int
	dict map[int]int
}

func (ts *TwoSum) AddInt(i int) {
	ts.list = append(ts.list, i)
	ts.dict[i] = len(ts.list) - 1
}

func (ts *TwoSum) Find(v int) bool{
	for i := range ts.list {
		if _, ok := ts.dict[v-ts.list[i]]; ok {
			return true
		} 
	}
	return false
}

func NewTwoSum() iTwoSum {
	return &TwoSum{ make([]int, 0, 16), make(map[int]int) }
}

func main() {
	ts := NewTwoSum()
	ts.AddInt(1)
	ts.AddInt(3)
	ts.AddInt(5)
	fmt.Println(ts.Find(4))
	fmt.Println(ts.Find(7))
}
