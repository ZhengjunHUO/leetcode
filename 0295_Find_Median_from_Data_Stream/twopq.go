package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

type MedianFinder struct {
	// 存放较小的一半数
	SmallPopMax	*godtype.PriorityQueue 
	// 存放较大的一半数
	BigPopMin	*godtype.PriorityQueue 
	IsEven		bool
}

func Constructor() MedianFinder {
	return MedianFinder{
		SmallPopMax:	godtype.NewPQ([]int{}, []int{}, false),
		BigPopMin:	godtype.NewPQ([]int{}, []int{}, true),
		IsEven:		true,
	}
}

// 确保len(较大的pq) == len(较小的pq) 或 == len(较小的pq)+1
func (this *MedianFinder) AddNum(num int)  {
	if this.IsEven {
		this.SmallPopMax.Push(num, num)
		temp := this.SmallPopMax.Pop()
		this.BigPopMin.Push(temp, temp.(int))
	}else{
		this.BigPopMin.Push(num, num)
		temp := this.BigPopMin.Pop()
		this.SmallPopMax.Push(temp, temp.(int))
	}
	
	this.IsEven = !this.IsEven
}


func (this *MedianFinder) FindMedian() float64 {
	if this.IsEven {
		return (float64(this.SmallPopMax.Peek().(int)) + float64(this.BigPopMin.Peek().(int))) / 2
	}else{
		return float64(this.BigPopMin.Peek().(int))
	}
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}
