package main

func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}

	if n%2 == 0 {
		return getMaximumGenerated(n/2)
	}

	return getMaximumGenerated(n/2) + getMaximumGenerated(n/2 + 1)
}
