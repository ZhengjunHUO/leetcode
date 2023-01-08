package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type LRUCache struct {
	Cap	int
	LMap	*datastruct.LinkedHashmap
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, datastruct.NewLinkedHashmap()}
}

func (this *LRUCache) Get(key int) int {
	if ! this.LMap.Contains(key) {
		return -1
	}

	// 将被访问的元素放到最后（最新）
	this.LMap.BecomeNewest(key)
	return this.LMap.Get(key).(int)
}

func (this *LRUCache) Put(key int, value int) {
	if this.LMap.Contains(key) {
		// 更新元素的值，需要将该元素移到最后
		this.LMap.BecomeNewest(key)
        }

	this.LMap.Put(key, value)

	// 超过容量，移除最老的元素（最前面的元素）
	if this.LMap.Size() > this.Cap {
		this.LMap.PopEldest()
	}
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(1, 10)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
