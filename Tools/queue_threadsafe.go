package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	Elems []interface{}
	Lock  sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{}
}

// Push to the queue's end
func (q *Queue) Push(elem interface{}) {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	q.Elems = append(q.Elems, elem)
}

// Pop from the queue's head
func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}

	q.Lock.Lock()
	defer q.Lock.Unlock()

	defer func(){
		q.Elems = q.Elems[1:]
	}()
	
	return q.Elems[0]
}

func (q *Queue) IsEmpty() bool {
	q.Lock.RLock()
	defer q.Lock.RUnlock()

	return len(q.Elems) == 0 
}

func (q *Queue) Size() int {
	q.Lock.RLock()
	defer q.Lock.RUnlock()

	return len(q.Elems)
}

func main() {
	s := NewQueue()
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
