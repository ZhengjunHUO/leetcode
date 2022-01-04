package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)

	// time为0时，返回所有index
	if time == 0 {
		rslt := make([]int, n)
		for i := range security {
			rslt[i] = i		
		}
		return rslt
	}

	// 使用monotonic stack从后向前计算当前index向后下一个大于当前值的index
	mono_asc := godtype.NewStack()
	table_asc := make([]int, n)

	for i:=n-1; i>=0; i-- {
		// stack非空时，移除栈顶小于当前值的元素
		for !mono_asc.IsEmpty() && security[mono_asc.Peek().(int)] < security[i] {
			mono_asc.Pop()
		}

		if !mono_asc.IsEmpty() {
			// 栈顶元素(下一个值大于当前值的元素)的index和当前index只差为1表示找到一个递增
			if mono_asc.Peek().(int) - i == 1 {
				table_asc[i] = table_asc[i+1] + 1
			}
		}

		// 把当前元素推入栈
		mono_asc.Push(i)
	}

	// 反向用相同的方向生成递减表
	mono_desc := godtype.NewStack()
	table_desc := make([]int, n)

	for i:=0; i<n; i++ {
		for !mono_desc.IsEmpty() && security[mono_desc.Peek().(int)] < security[i] {
			mono_desc.Pop()
		}

		if !mono_desc.IsEmpty() {
			if i - mono_desc.Peek().(int) == 1 {
				table_desc[i] = table_desc[i-1] + 1
			}
		}

		mono_desc.Push(i)
	}

	//fmt.Println(table_asc)
	//fmt.Println(table_desc)

	// 对于某一天，只有其在两个表中的值都>=time才算合格
	rslt := []int{}
	for i := time; i < n - time; i++ {
		if table_asc[i] >= time && table_desc[i] >= time {
			rslt = append(rslt, i)
		}
	}

	return rslt
}

func main() {
	security := [][]int{[]int{5,3,3,3,5,6,2}, []int{1,1,1,1,1}, []int{1,2,3,4,5,6}}
	time := []int{2,0,2}

	for i := range security {
		fmt.Println(goodDaysToRobBank(security[i], time[i]))
	}
}
