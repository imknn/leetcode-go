package bit_manipulation

/* 231. 2的幂 */
func isPowerOfTwo(n int) bool {
	// 100 011
	return n > 0 && (n&(n-1) == 0)
}
