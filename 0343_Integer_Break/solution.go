package main

// 当n>4时，f(n) => f(n-3)*3 能得到最大的乘积
func integerBreak(n int) int {
	if n == 2 {
		// 1 * 1
		return 1
	}

	if n == 3 {
		// 1 * 2
		return 2
	}

	rslt := 1
	// n == 4时取2*2=4, n > 4时需要拆分为3*(n-3) 
	for n > 4 {
		n -= 3
		rslt *= 3
	}

	return rslt*n
}
