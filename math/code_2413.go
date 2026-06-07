package math

/* 2413. 最小偶倍数 */
// 如果n是奇数，那么就是n*2，否则就是n
func smallestEvenMultiple(n int) int {
	if n&1 == 1 { // 奇数
		return n << 1
	}
	return n
}
