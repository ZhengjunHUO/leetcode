package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type LFUCache struct {
	// 存放key, value
	KV	map[interface{}]interface{}
	// 存放key, frequence
	KF	map[interface{}]int
	// 存放frequence, keys
	FK	map[int]*datastruct.LinkedHashmap
	// 全局最小Freq
	MinF	int
	Cap	int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		KV:	make(map[interface{}]interface{}),
		KF:	make(map[interface{}]int),
		FK:	make(map[int]*datastruct.LinkedHashmap),
		MinF:	0,
		Cap:	capacity,
	}
}

func (this *LFUCache) IncremFreq(key int) {
	f := this.KF[key]
	// KF表中的key加1
	this.KF[key]++

	// FK表中对应Freq关联的linked hashmap中移除key代表的结点
	this.FK[f].Delete(key)
	// 此时如果该linked hashmap为空则删除之
	if this.FK[f].Size() == 0 {
		delete(this.FK, f)
		// 如果最小Frequence恰好是f则需要加1
		if f == this.MinF {
			this.MinF++
		}
	}

	// 在FK表中对应Freq+1的linked hashmap中加入
	if _, ok := this.FK[f+1]; !ok {
		this.FK[f+1] = datastruct.NewLinkedHashmap()
	}
	this.FK[f+1].Put(key, nil)
}

func (this *LFUCache) DeleteLFU() {
	//FK表中MinF关联的linked hashmap中移除LFU结点
	lm := this.FK[this.MinF]
	keyToDel := lm.PopEldest().Key

	if lm.Size() == 0 {
		delete(this.FK, this.MinF)
	}

	// 同时update KV, KF
	delete(this.KV, keyToDel)
	delete(this.KF, keyToDel)
}

func (this *LFUCache) Get(key int) int {
	if _, ok := this.KV[key]; !ok {
		return -1
	}

	this.IncremFreq(key)
	return this.KV[key].(int)
}

func (this *LFUCache) Put(key int, value int)  {
	// 如果key已存在则更新对应的value，增加访问次数
	if _, ok := this.KV[key]; ok {
		this.KV[key] = value
		this.IncremFreq(key)
		return
	}

	// key不存在，需要加入
	// 如果cache已满则移除LFU
	if len(this.KV) == this.Cap {
		this.DeleteLFU()
	}

	// update KV, KF
	this.KV[key] = value
	this.KF[key] = 1

	// update FK
	if _, ok := this.FK[1]; !ok {
		this.FK[1] = datastruct.NewLinkedHashmap()
	}
	this.FK[1].Put(key, nil)

	this.MinF = 1
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
