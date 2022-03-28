package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// eg. if charToReplace is 'T', return the maximum number of consecutive 'F'
func maxOccuByReplacingChar(answerKey string, k int, charToReplace byte) int {
	// left, right: window's bondary
	// cnt: current number of replacement performed
	var left, cnt, ret int

	// right bondary continue to move on
	for right := range answerKey {
		// do replacement
		if answerKey[right] == charToReplace {
			cnt++
		}

		// number of operation exceed limit k, need to shrink the left boundry until == k
		for cnt > k {
			// if found a charToReplace on the left, decrement counter
			if answerKey[left] == charToReplace {
				cnt--
			}
			left++
		}

		// update the max to return
		ret = max(ret, right-left+1)
	}

	return ret
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxOccuByReplacingChar(answerKey, k, 'T'), maxOccuByReplacingChar(answerKey, k, 'F'))
}

func main() {
	answerKey := []string{"TTFF", "TFFT", "TTFTTFTT"}
	k := []int{2, 1, 1}
	for i := range k {
		fmt.Println(maxConsecutiveAnswers(answerKey[i], k[i]))
	}
}
