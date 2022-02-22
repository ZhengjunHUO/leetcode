package extra

func XorProductBF(M int, N int) int {
	sum := 0

	if M == N {
		return 0
	}

	for i:=M; i<=N; i++ {
		sum ^= i
	}

	return sum
}

