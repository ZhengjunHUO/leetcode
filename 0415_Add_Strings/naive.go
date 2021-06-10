package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	rslt := ""
	l1, l2, v1, v2, carry := len(num1)-1, len(num2)-1, 0, 0, 0
	
	for l1 >= 0 || l2 >= 0 || carry > 0 {
		if l1 >=0 {
			v1 = int(num1[l1] - '0')
			l1 -= 1
		}else{
			v1 = 0
		}

		if l2 >=0 {
			v2 = int(num2[l2] - '0')
			l2 -= 1
		}else{
			v2 = 0
		}

		sum := v1 + v2 + carry
		carry = sum/10

		rslt = fmt.Sprintf("%v%s", rune(sum%10), rslt)
	}

	return rslt    
}

func main() {
	fmt.Println(addStrings("11","123"))
	fmt.Println(addStrings("456","77"))
	fmt.Println(addStrings("0","0"))
}
