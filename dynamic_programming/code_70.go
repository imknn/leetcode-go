package dynamic_programming

/* 70. 爬楼梯 */
func climbStairs(n int) int {
	if n <= 2 {
		return n
	} else {
		a, b := 1, 2
		for ; n > 2; n-- {
			a, b = b, a+b
		}
		return b
	}
}
