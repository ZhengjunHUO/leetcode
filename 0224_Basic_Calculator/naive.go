package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func calculate(s string) int {
	return calc(&s)
}

func calc(s *string) int {
	currNum, sign := 0, byte('+')
	stack := godtype.NewStack()	
	
	for len(*s) > 0 {
		currByte := (*s)[0]
		*s = (*s)[1:]

		if currByte == byte('(') {
			currNum = calc(s)
		}

		if currByte >= byte('0') && currByte <= byte('9') {
			currNum = 10 * currNum + int(currByte - byte('0'))
		}

		if ( (currByte < byte('0') || currByte > byte('9')) && currByte != byte(' ') ) || len(*s) == 0 {
			if sign == byte('+') {
				stack.Push(currNum)
			}else{
				stack.Push(-currNum)
			}
			
			currNum = 0
			sign = currByte
		}

		if currByte == byte(')') {
			break
		}
	}

	rslt := 0
	for !stack.IsEmpty() {
		rslt += stack.Pop().(int)
	}

	return rslt
}

func main() {
	ss := []string{"1 + 1", " 2-1 + 2 ", "(1+(4+5+2)-3)+(6+8)", "((-3+5) + 8) - 4 "}
	for i := range ss {
		fmt.Println(calculate(ss[i]))
	}
}
