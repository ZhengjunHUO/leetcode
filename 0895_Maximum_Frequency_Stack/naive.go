package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

type FreqStack struct {
	// 存放值和该值的出现频率
	VF	map[int]int
	// 出现频率和存放出现过f次的值
	FV	map[int]*godtype.Stack
	// 当前最大的出现频率
	MaxF	int
}

func Constructor() FreqStack {
	return FreqStack{
		VF:	make(map[int]int),
		FV:	make(map[int]*godtype.Stack),
		MaxF:	0,
	}
}

func (this *FreqStack) Push(val int)  {
	// val的出现频率加1
	this.VF[val]++
	f := this.VF[val]

	if f > this.MaxF {
		this.MaxF = f
	}

	// 加入新频率对应的stack
	if _, ok := this.FV[f]; !ok {
		this.FV[f] = godtype.NewStack()
	}
	this.FV[f].Push(val)	
}

func (this *FreqStack) Pop() int {
	// 找到当前最大出现频率的stack，弹出最近插入的值
	s := this.FV[this.MaxF]
	rslt := s.Pop().(int)

	if s.Size() == 0 {
		delete(this.FV, this.MaxF)
		this.MaxF--

		/*
		for this.MaxF > 0 {
			if _, ok := this.FV[this.MaxF]; ok {
				break
			}
			this.MaxF--		
		}
		*/
	}

	this.VF[rslt]--

	return rslt
}

func main() {
	obj := Constructor()
	obj.Push(5)
	obj.Push(7)
	obj.Push(5)
	obj.Push(7)
	obj.Push(4)
	obj.Push(5)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
}
