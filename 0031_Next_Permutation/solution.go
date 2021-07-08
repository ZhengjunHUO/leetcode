package main

import (
	"fmt"
)

func findFirstDescIdx(nums []int) (int, int) {
	// 将遇到的元素存入hashmap，key为值，value存放元素index
	dict := make(map[int]int)

	for i:=len(nums)-1; i>0; i-- {
		// 如果元素值相同则只存第一次出现的index
		// 目的是如果后续需要交换，则能确保交换后的数列还是递减数列
		if _, ok := dict[nums[i]]; !ok {
			dict[nums[i]] = i
		}

		// 遇到（从后往前的）第一个递减元素
		if nums[i-1] < nums[i] {
			firstDesc := i-1 

			// 向后寻找刚刚好大于它的元素index
			diff, nextBigger := 101, firstDesc
			for k,v := range dict {
				// 注意判断条件
				if temp:= k - nums[firstDesc]; temp < diff && temp > 0 {
					diff = temp
					nextBigger = v
				}				
			}
			
			return firstDesc, nextBigger
		}
	}

	return -1, -1
}


func nextPermutation(nums []int)  {
	// 从后往前寻找第一个递减元素的index，然后再向后寻找刚刚好大于它的元素index
	firstDesc, nextBigger := findFirstDescIdx(nums)
	
	// 如果有找到符合条件的元素，将它们的值互换
	if firstDesc >= 0 {
		nums[firstDesc], nums[nextBigger] = nums[nextBigger], nums[firstDesc]
	}
    
	// 将递减元素index（值已经互换）以后的数列倒序排列（即递减数列变为递增数列）
	// firstDesc为-1的情况下即原数列是个递减数列，没有下一个permutation，倒序排列后正好符合题目要求
	lp, rp := firstDesc + 1, len(nums) - 1
	for lp < rp {
		nums[lp], nums[rp] = nums[rp], nums[lp]
		lp += 1
		rp -= 1
	} 
}

func main() {
	arraySet := [][]int{[]int{1,2,3}, []int{3,2,1}, []int{1,1,5}, []int{1}, []int{1,5,8,4,7,6,6,3,1}}
	for _, v := range arraySet {
		nextPermutation(v)
		fmt.Println(v)
	}
}
