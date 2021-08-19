package main

import (
	"fmt"
	"math/rand"
)

// 不能只用hashmap的原因在于需要O(1)均匀随机得返回存储的值
// 配合一段连续的存储空间可以达成目标
type RandomizedSet struct {
	// 映射值和index
	VI	map[int]int    
	Elems	[]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), []int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.VI[val]; ok {
		return false
	}

	this.VI[val] = len(this.Elems)
	this.Elems = append(this.Elems, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.VI[val]; !ok {
		return false
	}
	
	idx := this.VI[val]
	this.VI[this.Elems[len(this.Elems)-1]] = idx
	delete(this.VI, val)

	this.Elems[idx], this.Elems[len(this.Elems)-1] = this.Elems[len(this.Elems)-1], this.Elems[idx]	
	this.Elems[len(this.Elems)-1] = 0
	this.Elems = this.Elems[:len(this.Elems)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.Elems[rand.Int()%len(this.Elems)]
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.Insert(2))
	for i:=0; i<10; i++ {
		fmt.Printf("%v ", obj.GetRandom())
	}
	fmt.Println()
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Insert(3))
	for i:=0; i<10; i++ {
		fmt.Printf("%v ", obj.GetRandom())
	}
	fmt.Println()
}
