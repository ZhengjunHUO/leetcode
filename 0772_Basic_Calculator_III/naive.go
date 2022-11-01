package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

// 同0227_Basic_Calculator_II
func calculate(s string) int {
	return calc(&s)
}

func calc(s *string) int {
	currNum, sign := 0, byte('+')
	stack := datastruct.NewStack([]int{})

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
			switch sign {
			case byte('+'):
				stack.Push(currNum)
			case byte('-'):
				stack.Push(-currNum)
			case byte('*'):
				stack.Push(currNum * stack.Pop())
			case byte('/'):
				stack.Push(stack.Pop() / currNum)
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
		rslt += stack.Pop()
	}

	return rslt
}

func main() {
	ss := []string{"1 + 1", " 6-4 / 2 ", "2*(5+5*2)/3+(6/2+8)", "(2+6* 3+5- (3*14/7+2)*5)+3"}
	for i := range ss {
		fmt.Println(calculate(ss[i]))
	}
}
