package math

/* 258. 各位相加 */

func addDigits(num int) int {
	m, count := num, 0
	for m >= 10 {
		count++
		m = 0
		for num > 0 {
			m += num % 10
			num /= 10
		}
		num = m
	}
	return m
}

//func addDigits(num int) int {
//	return (num-1)%9 + 1
//}
