package bit_manipulation

/* 461. 汉明距离 */
func hammingDistance(x int, y int) int {
	m := x ^ y // 异或计算不同的位置
	count := 0
	for m > 0 { // 统计1的个数
		if m&1 > 0 {
			count++
		}
		m >>= 1
	}
	return count
}
