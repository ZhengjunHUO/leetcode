package main

import (
	"fmt"
	"sync"
)

type Deque struct {
	Elems []interface{}
	Lock  sync.RWMutex
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) pushFirst(elem interface{}) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elems = append([]interface{}{elem}, d.Elems...)
}

func (d *Deque) pushLast(elem interface{}) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elems = append(d.Elems, elem)
}

func (d *Deque) popFirst() interface{} {
	if d.isEmpty() {
		return nil
	}

	d.Lock.Lock()
	defer d.Lock.Unlock()

	defer func(){
		d.Elems = d.Elems[1:]
	}()
	
	return d.Elems[0]
}

func (d *Deque) popLast() interface{} {
	if d.isEmpty() {
		return nil
	}

	d.Lock.Lock()
	defer d.Lock.Unlock()

	n := len(d.Elems)

	defer func(){
		d.Elems = d.Elems[:n-1]
	}()
	
	return d.Elems[n-1]
}

func (d *Deque) peekFirst() interface{} {
	if d.isEmpty() {
		return nil
	}

	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elems[0]
}

func (d *Deque) peekLast() interface{} {
	if d.isEmpty() {
		return nil
	}

	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elems[len(d.Elems)-1]
}

func (d *Deque) isEmpty() bool {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elems) == 0 
}

func (d *Deque) Size() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elems)
}


func main() {
	d := NewDeque()
	d.popFirst()
	d.popLast()
	d.pushFirst(3)
	d.pushLast(5)
	d.pushFirst(2)
	d.pushLast(10)
	fmt.Println(d.Elems)
	fmt.Println(d.Size())
	fmt.Println(d.peekFirst())
	fmt.Println(d.popFirst())
	fmt.Println(d.popLast())
	fmt.Println(d.Elems)
}
