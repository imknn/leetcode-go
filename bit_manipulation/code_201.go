package bit_manipulation

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	num := m
	for i := m + 1; i <= n; i++ {
		num &= i
	}
	return num
}
