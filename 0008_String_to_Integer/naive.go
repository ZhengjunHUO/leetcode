package main

import (
	"fmt"
)

const MAXVAL int = 2147483647
const MINVAL int = -2147483648
const MAXVALDIVTEN int = 214748364
const ASCIIZERO uint8 = 48
const ASCIISEPT uint8 = 55
const ASCIIHUIT uint8 = 56
const ASCIININE uint8 = 57

func myAtoi(s string) int {
	n := len(s) 
	if n == 0 {
		return 0
	}

	i, sign, rslt := 0, 1, 0

	// 去除开始的空格
	for i < n {
		if s[i] == ' ' {
			i+=1
		}else{
			break
		}
	}

	// 处理可能存在的符号
	if s[i] == '+' || s[i] == '-' {
		if i == n - 1 {
			return rslt
		}

		if s[i] == '-' {
			sign = -1
		}

		i+=1
	}

	for i < n {
		if s[i] >= ASCIIZERO && s[i] <= ASCIININE {
			rslt *= 10
			// 判断已读取的数后面append最新的一位是否会超过最大值
			if (rslt > MAXVALDIVTEN) || (rslt == MAXVALDIVTEN && ( (s[i] > ASCIISEPT && sign>0) || (s[i] > ASCIIHUIT && sign<0))){
				if sign > 0 {
					return MAXVAL
				}else{
					return MINVAL
				}
			}

			rslt += int(s[i]-ASCIIZERO)
			i++
		}else{
			return rslt * sign
		}
	}	
    
	return rslt * sign
}

func main() {
	str := []string{"42", "   -42", "4193 with words", "words and 987", "-91283472332", "2147483647", "2147483648", "-2147483648", "-2147483649" }
	for _, s := range str {
		fmt.Printf("Input: %s\t Output: %d\n", s, myAtoi(s))
	}
}
