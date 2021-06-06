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
	return len(q.Elems)
}

func minJumps(arr []int) int {
	n, jmpCnt := len(arr), 0
	if n <= 1 {
		return 0
	}

	posDict := make(map[int][]int)
	for i := range arr {
		posDict[arr[i]] = append(posDict[arr[i]], i)
	}    
//	fmt.Println(posDict)
	
	s := NewQueue()
	s.Push(0)
//	fmt.Println(s.Elems)
	
	for !s.IsEmpty() {
		jmpCnt += 1
		currLvlSize := s.Size()

		for i:=0; i<currLvlSize; i++ {
			idx := s.Pop().(int)

			if v, ok := posDict[arr[idx]]; ok {
//				fmt.Println("v: ", v)
				for _,j := range v {
					if idx != j {
						if j == n - 1 {
							return jmpCnt
						}
						s.Push(j)
//						fmt.Println("Push j: ", j)
					}
	
				}
				delete(posDict, arr[idx])
//				fmt.Println(posDict)
			} 

			if idx - 1 > 0 {
				if _, ok := posDict[arr[idx-1]]; ok {
					s.Push(idx-1)		
//					fmt.Println("Push idx-1: ", idx-1)
				}
			}

			if idx + 1 < n {
				if _, ok := posDict[arr[idx+1]]; ok {
					if idx+1 == n-1 {
						return jmpCnt
					}

					s.Push(idx+1)
//					fmt.Println("Push idx+1: ", idx+1)
				}
			}

//			fmt.Println(s.Elems)
		}
	}	

	return jmpCnt
}

func main() {
	arr := []int{100,-23,-23,404,100,23,23,23,3,404}
	fmt.Println(minJumps(arr))
//	fmt.Println("OK: ", minJumps(arr))

	arr = []int{7}
	fmt.Println(minJumps(arr))
//	fmt.Println("OK: ", minJumps(arr))

	arr = []int{7,6,9,6,9,6,9,7}
	fmt.Println(minJumps(arr))
//	fmt.Println("OK: ", minJumps(arr))

	arr = []int{6,1,9}
	fmt.Println(minJumps(arr))
//	fmt.Println("OK: ", minJumps(arr))

	arr = []int{11,22,7,7,7,7,7,7,7,22,13}
	fmt.Println(minJumps(arr))
//	fmt.Println("OK: ", minJumps(arr))
}
