package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	rslt := ""
	l1, l2, v1, v2, carry := len(a)-1, len(b)-1, 0, 0, 0

	for l1 >=0 || l2 >= 0 || carry > 0 {
		if l1 >= 0 {
			v1 = int(a[l1] - '0')
			l1 -= 1
		}else{
			v1 = 0
		}
		if l2 >= 0 {
			v2 = int(b[l2] - '0')
			l2 -= 1
		}else{
			v2 = 0
		}

		sum := v1 + v2 + carry
		carry = sum/2
		rslt = fmt.Sprintf("%v%s", sum%2, rslt)
	}

	return rslt    
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
