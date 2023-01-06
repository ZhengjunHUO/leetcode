package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type MedianFinder struct {
	// 存放较小的一半数
	SmallPopMax	*datastruct.PriorityQueue[int, int]
	// 存放较大的一半数
	BigPopMin	*datastruct.PriorityQueue[int, int]
	IsEven		bool
}

func Constructor() MedianFinder {
	return MedianFinder{
		SmallPopMax:	datastruct.NewPQ([]int{}, []int{}, false),
		BigPopMin:	datastruct.NewPQ([]int{}, []int{}, true),
		IsEven:		true,
	}
}

// 确保len(较大的pq) == len(较小的pq) 或 == len(较小的pq)+1
func (this *MedianFinder) AddNum(num int)  {
	if this.IsEven {
		this.SmallPopMax.Push(num, num)
		temp := this.SmallPopMax.Pop()
		this.BigPopMin.Push(temp, temp)
	}else{
		this.BigPopMin.Push(num, num)
		temp := this.BigPopMin.Pop()
		this.SmallPopMax.Push(temp, temp)
	}

	this.IsEven = !this.IsEven
}


func (this *MedianFinder) FindMedian() float64 {
	if this.IsEven {
		return (float64(this.SmallPopMax.Peek()) + float64(this.BigPopMin.Peek())) / 2
	}else{
		return float64(this.BigPopMin.Peek())
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
