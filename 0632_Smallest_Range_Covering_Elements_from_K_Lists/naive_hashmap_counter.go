package main

import (
	"fmt"
)

func smallestRange(nums [][]int) []int {
	n := len(nums)

	//保存各子串当前的index，未来遍历用
	pointers := make([]int, n)
	//减为0 => valid状态（当前窗口涵盖了所有子串中至少一个元素）
	counter := n
	//增为n表明已经遍历所有子串，可以退出了
	endedList := 0

	//初始化字典，1为需要1个，0为该子串已经valid，-1表示多包含了1个。。。
	m := make(map[int]int)
	for i:=0; i<n; i++ {
		m[i] = 1
	}

	//所有子串元素按顺序整合后的列表，内容为（“来源”，“值”）
	mergedList := [][2]int{}
	lp, rp := 0, 0
	start, minRange := -1, 200000
	
	// 可以改为for true
	for endedList < n {
		endedList = 0
		candidateList := []int{}

		// 搜集所有子串最小值做成列表
		for i:=0; i<n; i++ {
			//该子串已经遍历完成，但仍旧需要一个最大值占位
			if pointers[i] == len(nums[i]) {
				endedList += 1
				candidateList = append(candidateList, 100000)
			}else{
				candidateList = append(candidateList,nums[i][pointers[i]])
			}
		}

		//已经遍历所有子串，退出
		if endedList == n {
			break
		}
		
		// 获取所有子串最小值中的最小值，返回的index为第几个子串
		idx := minIdx(candidateList, 0, n-1)
		
		if (m[idx]) > 0 {
			counter -= 1
		}
		m[idx] -= 1
		mergedList = append(mergedList, [2]int{idx, nums[idx][pointers[idx]]})
		pointers[idx] += 1
		// 每一步都移动右指针 
		rp += 1

		// 移动左指针直到退出valid状态
		for (counter == 0) {
			if temp := mergedList[rp-1][1] - mergedList[lp][1]; temp < minRange {
				start = lp
				minRange = temp
			}	

			m[mergedList[lp][0]] += 1
			if (m[mergedList[lp][0]]) > 0 {
				counter += 1
			}
			lp += 1
		}

	}
	
	if minRange < 200000 {
		return []int{mergedList[start][1], mergedList[start][1]+minRange}
	}

	return []int{}
}

func minIdx(l []int, lp int, rp int) int {
	if lp == rp {
		return lp
	}
	if rp - lp == 1 {
		if l[lp] <= l[rp] {
			return lp 
		}else{
			return rp
		}
	}

	mid := lp + (rp - lp) / 2
	lm := minIdx(l, lp, mid)
	rm := minIdx(l,mid+1,rp)
	if l[lm] <= l[rm] {
		return lm
	}else{
		return rm
	}
}

func main() {
	nums := [][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}}
	fmt.Println(smallestRange(nums))

	nums = [][]int{{1,2,3},{1,2,3},{1,2,3}}
	fmt.Println(smallestRange(nums))

	nums = [][]int{{10,10},{11,11}}
	fmt.Println(smallestRange(nums))

	nums = [][]int{{1},{2},{3},{4},{5},{6},{7}}
	fmt.Println(smallestRange(nums))
}
