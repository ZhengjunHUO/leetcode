package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	Elems	[]interface{}
	Lock	sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return len(s.Elems) == 0
}

func (s *Stack) Push(elem interface{}) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Elems = append(s.Elems, elem)	
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}	

	s.Lock.Lock()
	defer s.Lock.Unlock()

	n := len(s.Elems)

	defer func(){
		s.Elems = s.Elems[:n-1]
	}() 

	return s.Elems[n-1] 
}

func (s *Stack) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return len(s.Elems)
}

func main() {
	s := NewStack()
	fmt.Println(s.Pop())

	s.Push(3)
	s.Push(6)
	s.Push(9)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
}
