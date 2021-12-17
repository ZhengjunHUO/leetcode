package extra

/* 
  do a XOR production from 0 we can observe a rule of cycle of 4 
	x  Bin  XOR_Prod_from_0
	0  0000 0000 => 0 [x]
	1  0001 0001 => 1 [1]
	2  0010 0011 => 3 [x+1]
	3  0011 0000 => 0 [0]
	4  0100 0100 => 4 [x]
	5  0101 0001 => 1 [1]
	6  0110 0111 => 7 [x+1]
	7  0111 0000 => 0 [0]
	8  1000 1000 => 8 [x]
	9  1001 0001 => 1 [1]
	10 1010 1011 => 11 [x+1]
	11 1011 0000 => 0 [0]
*/
func prodFromZeroToX(x int) int {
	prod := []int{x, 1, x+1, 0}
	return prod[x%4]
}

/*
   because:
	wanted = M ^ (M+1) ^ ... ^ (N-1) ^ N
	prodFromZeroToX(N) = (0 ^ 1 ^ ... ^ M-1) ^ (M ^ M+1 ^ ... ^ N-1 ^ N)
			   = prodFromZeroToX(M-1) ^ wanted
   then XOR prodFromZeroToX(M-1) in both side we have:
	wanted = prodFromZeroToX(N) ^ prodFromZeroToX(M-1)
*/
func XorProduct(M int, N int) int {
	if M == N {
		return 0
	}

	return prodFromZeroToX(N) ^ prodFromZeroToX(M-1)
}

