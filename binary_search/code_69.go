package binary_search

/* 69. x 的平方根 */
func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	if x < 4 {
		return 1
	}
	l, r, m := 1, x, x/2
	for {
		if x/m < m {
			r = m
			m = (l + r) / 2
			if r == m {
				return m
			}
			continue
		}
		if x/m > m {
			l = m
			m = (l + r) / 2
			if l == m {
				return m
			}
			continue
		}
		return m
	}
}
