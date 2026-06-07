package math

/* 728. 自除数 */
func selfDividingNumbers(left int, right int) []int {
	m := make([]int, 0)
	for i := left; i <= right; i++ {
		a := i
		for a > 0 {
			b := a % 10
			if b == 0 {
				break
			} else if b != 1 && i%b != 0 {
				break
			}
			a /= 10
		}
		if a == 0 {
			m = append(m, i)
		}
	}
	return m
}
