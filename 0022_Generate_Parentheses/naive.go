package main

import (
	"fmt"
)

func backtrack(rslt *[]string, target int, curr string, opened int, completed int) {
	// 打开的左括号数量加上已成对的左括号数量大于目标数，直接退出
	if opened + completed > target {
		return
	}

	if completed == target {
		// 找到target对括号(此时opened只能为0)
		*rslt = append(*rslt, curr)
		return 
	}

	if opened == 0 {
		backtrack(rslt, target, curr+"(", 1, completed)
	}else{
		backtrack(rslt, target, curr+")", opened-1, completed+1)
		backtrack(rslt, target, curr+"(", opened+1, completed)
	}
}

func generateParenthesis(n int) []string {
	rslt := []string{}
	backtrack(&rslt, n, "(", 1, 0)
	return rslt    
}

func main() {
	ints := []int{3, 1}
	for _, v := range ints {
		fmt.Println(generateParenthesis(v))
	}
}
